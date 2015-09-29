package spawnr

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRun(t *testing.T) {
	// fmt.Println("@testGoRun()")
	// defer KillAll()
	// go GoRun("test1", 2)
	// go GoRun("test2", 2)
	// time.Sleep(100 * time.Millisecond)

	// //	pid := 445
	// // cmd := "ps -g "+string(pid)+" | grep -v "+string(pid)
	// // ps, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	// // fmt.Printf("%s\n", ps)
	// fmt.Println("Sleeping")
	// time.Sleep(1000 * time.Millisecond)
	// fmt.Println("Exiting")
	// t.Fail()
}

func TestExec(t *testing.T) {
	fmt.Println("@testExec()")

	defer KillAll()
	go Exec("redis-server", 2)
	time.Sleep(30 * time.Second)

	fmt.Println("Exiting")
	t.Fail()
}
