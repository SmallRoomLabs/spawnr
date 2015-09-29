package spawnr

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const (
	RUNNING = 1
	STOPPED = 2
)

const (
	EXEC  = 1
	GORUN = 2
)

type App struct {
	apptype  int
	killtype int
	status   int
	name     string
	pid      int
}

var runningApps = make(map[int]App)

//
// Helper function for modifying a field in a map
//
func (t App) setStatus(st int) App {
	t.status = st
	return t
}

//
//
//
func KillAll() {
	fmt.Println("@KillAllApps()")

	for _, v := range runningApps {
		if v.status == RUNNING {
			pid := v.pid
			if v.apptype == EXEC {
				log.Printf("Trying to kill %s with PID %d \n", v.name, pid)
				if v.killtype == 1 {
					log.Printf("SIGKILL(%d)\n", pid)
					syscall.Kill(pid, syscall.SIGKILL)
				}
				if v.killtype == 2 {
					log.Printf("SIGTERM(%d)\n", pid)
					syscall.Kill(pid, syscall.SIGTERM)
				}
			}
			if v.apptype == GORUN {
				log.Printf("Trying to kill %s with PID %d and its siblings\n", v.name, pid)
				cmd := "ps -f | cut -c 7-19 | grep " + strconv.Itoa(pid) + " | cut -c 1-5"
				plist, _ := exec.Command("/bin/bash", "-c", cmd).Output()
				for _, p := range strings.Split(string(plist), "\n") {
					pn, _ := strconv.Atoi(p)
					if pn > 0 {
						if v.killtype == 1 {
							log.Printf("SIGKILL(%d)\n", pn)
							syscall.Kill(pn, syscall.SIGKILL)
						}
						if v.killtype == 2 {
							log.Printf("SIGTERM(%d)\n", pn)
							syscall.Kill(pn, syscall.SIGTERM)
						}
					}
				}
			}
		}
	}
}

//
//
//
func GoRun(app string, killtype int) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("/usr/local/go/bin/go", "run", app+"/"+app+".go")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	pid := cmd.Process.Pid

	runningApps[pid] = App{apptype: GORUN, killtype: killtype, status: RUNNING, name: app, pid: pid}
	log.Printf("Running %s in PID %d\n", app, pid)

	cmd.Wait()
	log.Printf("PID %d finished with stdout: %s\n", pid, stdout.String())
	log.Printf("PID %d finished with stderr: %s\n", pid, stderr.String())
	runningApps[pid] = runningApps[pid].setStatus(STOPPED)
}

//
//
//
func Exec(app string, killtype int) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(app)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	pid := cmd.Process.Pid

	runningApps[pid] = App{apptype: EXEC, killtype: killtype, status: RUNNING, name: app, pid: pid}
	log.Printf("Running %s in PID %d\n", app, pid)

	cmd.Wait()
	log.Printf("PID %d finished with stdout: %s\n", pid, stdout.String())
	log.Printf("PID %d finished with stderr: %s\n", pid, stderr.String())
	runningApps[pid] = runningApps[pid].setStatus(STOPPED)
}
