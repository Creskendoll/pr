package main

import (
	"os"
	"os/exec"
)

func gh(args ...string) (string, error) {
	cmd := exec.Command("gh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "", nil
}
