package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: timer.exe <command> [args]")
		os.Exit(1)
	}

	cmdArgs := os.Args[1:]

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil

	start := time.Now()

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command: ", err)
		os.Exit(1)
	}

	duration := time.Since(start)

	fmt.Printf("%v\n", duration)
}
