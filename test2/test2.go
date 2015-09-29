package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("TEST2\n")
	time.Sleep(15 * time.Second)
	os.Exit(2)
}
