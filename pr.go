package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	ctx := context.Background()

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

	llmClient, err := LLMClient(ctx)
	if err != nil {
		fmt.Println("Error creating LLM client:", err)
		os.Exit(1)
	}

	description, err := diffDescription(llmClient, diff, ctx)
	if err != nil {
		fmt.Println("Error getting diff description:", err)
		os.Exit(1)
	}

	tmpFilePath := filepath.Join(os.TempDir(), fmt.Sprintf("pr-%s.md", gitBranch))
	err = os.WriteFile(tmpFilePath, []byte(description), 0755)
	if err != nil {
		fmt.Println("Error writing description to file:", err)
		os.Exit(1)
	}

	err = exec.Command("open", tmpFilePath).Run()
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	fmt.Println("Does the description look good? (y/n)")
	var answer string
	n, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	if n != 1 || strings.ToLower(answer) != "y" || strings.ToLower(answer) != "yes" {
		fmt.Println("Exiting")
		os.Exit(0)
	}

	editedDescription, err := os.ReadFile(tmpFilePath)
	if err != nil {
		fmt.Println("Error reading edited description:", err)
		os.Exit(1)
	}

	prCreated, err := gh("pr", "create", "--title", fmt.Sprintf("\"%s\"", gitBranch), "--body", fmt.Sprintf("\"%s\"", editedDescription))
	if err != nil {
		fmt.Println("Error creating PR:", err)
		os.Exit(1)
	}
	fmt.Println("PR created:", prCreated)
}
