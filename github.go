package main

import (
	"os"
	"os/exec"

	"github.com/google/go-github/v69/github"
)

func getToken() (string, error) {
	ghEnv := os.Getenv("PR_GITHUB_KEY")

	if ghEnv != "" {
		return ghEnv, nil
	}

	getTokenCommand := exec.Command("gh", "auth", "token")
	token, err := getTokenCommand.Output()
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func GitHubClient() (*github.Client, error) {
	ghToken, err := getToken()
	if err != nil {
		return nil, err
	}
	client := github.NewClient(nil).WithAuthToken(ghToken)
	return client, nil
}
