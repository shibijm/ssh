package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	sshBinPath := os.Getenv("SSH_BIN_PATH")
	if sshBinPath == "" {
		fmt.Println("Environment variable SSH_BIN_PATH is not set")
		os.Exit(1)
	}
	signalChannel := make(chan os.Signal, 1)
	for {
		// err := initSshAgent()
		// if err != nil {
		// 	panic(err)
		// }
		cmd := exec.Command(sshBinPath, os.Args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		signal.Notify(signalChannel, os.Interrupt)
		err := cmd.Run()
		signal.Stop(signalChannel)
		if err == nil || len(os.Args) == 1 {
			break
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ProcessState.ExitCode() == 512 {
				break
			}
		}
		fmt.Printf("Reconnecting... (%s)\n", err)
		time.Sleep(1 * time.Second)
	}
}
