package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

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

	diff, err := getDiff()
	if err != nil {
		fmt.Println("Error getting git diff:", err)
		os.Exit(1)
	}

	description, err := diffDescription(llmClient, diff, ctx)
	if err != nil {
		fmt.Println("Error getting diff description:", err)
		os.Exit(1)
	}

	prCreated, err := gh("pr", "create", "--title", fmt.Sprintf("\"%s\"", gitBranch), "--body", fmt.Sprintf("\"%s\"", description))
	if err != nil {
		fmt.Println("Error creating PR:", err)
		os.Exit(1)
	}
	fmt.Println("PR created:", prCreated)
}
