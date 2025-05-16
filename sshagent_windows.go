package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

const WINDOWS_PROC_CREATE_NEW_PROCESS_GROUP = 0x00000200
const WINDOWS_PROC_DETACHED_PROCESS = 0x00000008

func initSshAgent() error {
	tempDir := os.TempDir()
	entries, err := os.ReadDir(tempDir)
	if err != nil {
		return err
	}
	cmd := exec.Command("tasklist", "/FO", "CSV", "/NH", "/FI", "IMAGENAME eq ssh-agent.exe")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	output := string(outputBytes)
	if strings.Contains(output, "No tasks are running") {
		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), "ssh-") {
				sshAgentTempDir := filepath.Join(tempDir, entry.Name())
				err := os.RemoveAll(sshAgentTempDir)
				if err != nil {
					return err
				}
			}
		}
		fmt.Println("Starting ssh-agent")
		cmd := exec.Command("ssh-agent")
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: WINDOWS_PROC_DETACHED_PROCESS | WINDOWS_PROC_CREATE_NEW_PROCESS_GROUP}
		outputBytes, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		output := string(outputBytes)
		os.Setenv("SSH_AUTH_SOCK", regexp.MustCompile("SSH_AUTH_SOCK=(.*?);").FindStringSubmatch(output)[1])
		cmd = exec.Command("ssh-add")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return err
		}
		return nil
	} else {
		for _, entry := range entries {
			if strings.HasPrefix(entry.Name(), "ssh-") {
				sshAgentTempDir := filepath.Join(tempDir, entry.Name())
				subentries, err := os.ReadDir(sshAgentTempDir)
				if err != nil {
					return err
				}
				if len(subentries) > 0 {
					os.Setenv("SSH_AUTH_SOCK", filepath.Join(sshAgentTempDir, subentries[0].Name()))
					return nil
				}
			}
		}
		fmt.Println("Killing ssh-agent")
		cmd := exec.Command("taskkill", "/F", "/IM", "ssh-agent.exe")
		err := cmd.Run()
		if err != nil {
			return err
		}
		return initSshAgent()
	}
}
