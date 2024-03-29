package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]
	//
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if exitError, ok := err.(*exec.ExitError); ok {
		fmt.Println(exitError)
		os.Exit(exitError.ExitCode())
	}

	if err != nil {
		log.Fatal(err)
	}
}
