package main

import (
	"fmt"
	"os/exec"
)

func getDiff() (string, error) {
	remote, err := exec.Command("git", "remote").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get remote: %v", err)
	}

	defaultBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", fmt.Sprintf("%s/HEAD", remote)).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get default branch: %v", err)
	}

	baseCommit, err := exec.Command("git", "merge-base", string(defaultBranch), "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get base commit: %v", err)
	}

	diff, err := exec.Command("git", "diff", string(baseCommit)).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff: %v", err)
	}

	return string(diff), nil
}

func getBranch() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get branch: %v", err)
	}

	return string(output), nil
}
