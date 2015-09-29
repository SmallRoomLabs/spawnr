package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("TEST1\n")
	time.Sleep(10 * time.Second)
	os.Exit(0)
}
