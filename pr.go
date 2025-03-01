package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	githubClient, err := GitHubClient()
	if err != nil {
		fmt.Println("Error creating GitHub client:", err)
		os.Exit(1)
	}

	jiraClient, err := JiraClient()
	if err != nil {
		fmt.Println("Error creating Jira client:", err)
		os.Exit(1)
	}

	llmClient, err := LLMClient(ctx)
	if err != nil {
		fmt.Println("Error creating LLM client:", err)
		os.Exit(1)
	}

	gitBranch, err := getBranch()
	if err != nil {
		fmt.Println("Error getting git branch:", err)
		os.Exit(1)
	}
	fmt.Println("Current branch:", gitBranch)

	if githubClient != nil {
		fmt.Println("GitHub client created")
	}
	if jiraClient != nil {
		fmt.Println("Jira client created")
	}
	if llmClient != nil {
		fmt.Println("LLM client created")
	}
}
