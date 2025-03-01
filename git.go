package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func getDiff() (string, error) {
	remote, err := git("remote")
	if err != nil {
		return "", err
	}

	remoteOutput, err := git("remote", "show", remote)
	if err != nil {
		return "", err
	}

	headBranch, err := parseHeadBranch(remoteOutput)
	if err != nil {
		return "", err
	}

	baseCommit, err := git("merge-base", headBranch, "HEAD")
	if err != nil {
		return "", err
	}

	diff, err := git("diff", baseCommit)
	if err != nil {
		return "", err
	}

	return string(diff), nil
}

func getBranch() (string, error) {
	output, err := git("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}

	return output, nil
}

func parseHeadBranch(gitRemoteShowOutput string) (string, error) {
	re := regexp.MustCompile(`HEAD branch:\s+(\S+)`)
	matches := re.FindStringSubmatch(gitRemoteShowOutput)
	if len(matches) < 2 {
		return "", fmt.Errorf("failed to parse HEAD branch")
	}
	return strings.TrimSpace(matches[1]), nil
}

func git(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run git %v: %v", args, err)
	}
	return strings.TrimSpace(string(output)), nil
}
