package main

import (
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira"
)

func getJiraToken() (string, error) {
	jiraToken := os.Getenv("PR_JIRA_TOKEN")

	if jiraToken == "" {
		return "", fmt.Errorf("PR_JIRA_TOKEN is not set")
	}

	return jiraToken, nil
}

func JiraClient() (*jira.Client, error) {
	jiraToken, err := getJiraToken()

	if err != nil {
		return nil, err
	}

	auth := jira.BearerAuthTransport{
		Token: jiraToken,
	}
	jiraClient, err := jira.NewClient(auth.Client(), "https://issues.apache.org/jira/")

	if err != nil {
		return nil, err
	}

	return jiraClient, nil
}
