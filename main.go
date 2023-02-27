package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

func main() {
	var sshBinaryPath string
	if runtime.GOOS == "windows" {
		sshBinaryPath = `C:\Program Files\Git\usr\bin\ssh.exe`
	} else {
		sshBinaryPath = "/usr/bin/ssh"
	}
	if len(os.Args) > 1 && !strings.Contains(os.Args[1], "@") {
		fmt.Print("Username: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			username := strings.TrimSpace(scanner.Text())
			if len(username) == 0 {
				os.Exit(1)
			}
			os.Args[1] = username + "@" + os.Args[1]
		} else {
			fmt.Println()
			os.Exit(1)
		}
	}
	signalChannel := make(chan os.Signal)
	for {
		cmd := exec.Command(sshBinaryPath, os.Args[1:]...)
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
