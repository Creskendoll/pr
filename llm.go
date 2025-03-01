package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	llmApi "github.com/ollama/ollama/api"
)

func LLMClient(ctx context.Context) (*llmApi.Client, error) {
	llmClient, err := llmApi.ClientFromEnvironment()
	if err != nil {
		return nil, fmt.Errorf("failed to create LLM client: %v", err)
	}

	models, err := llmClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %v", err)
	}

	fmt.Println("Available models:")
	for _, model := range models.Models {
		fmt.Println(model.Name)
	}

	if len(models.Models) == 0 {
		supportedModels := []string{"DeepSeek-R1", "Llama 3.2", "Mistral", "Phi 4"}

		fmt.Println("No models found. Please select a model from the following list:")
		fmt.Println("0. Cancel")

		for i, model := range supportedModels {
			fmt.Printf("%d. %s\n", i+1, model)
		}

		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := consoleReader.ReadByte()
		modelIndex := int(input - '0')

		// ESC = 27 and Ctrl-C = 3
		if input == 27 || input == 3 || modelIndex == 0 {
			return nil, fmt.Errorf("setup cancelled")
		}

		if modelIndex > len(supportedModels) {
			return nil, fmt.Errorf("invalid model index")
		}
		model := supportedModels[modelIndex-1]

		fmt.Println("Pulling model:", model)
		err = pullModel(llmClient, model, ctx)
		if err != nil {
			return nil, err
		}
	}

	return llmClient, nil
}

func pullModel(llmClient *llmApi.Client, model string, ctx context.Context) error {
	req := &llmApi.PullRequest{
		Model: model,
		Name:  model,
	}

	bytesToGB := float32(1024.0 * 1024.0 * 1024.0)
	progressFunc := func(resp llmApi.ProgressResponse) error {
		if resp.Total > 0 {
			total := float32(resp.Total)
			completed := float32(resp.Completed)
			fmt.Print("\033[G\033[K")
			fmt.Printf("Downloading %.2f GB / %.2f GB (%.2f%%)\n", completed/bytesToGB, total/bytesToGB, completed/total*100.0)
			fmt.Print("\033[A")
		}
		return nil
	}

	err := llmClient.Pull(ctx, req, progressFunc)
	if err != nil {
		return fmt.Errorf("failed to pull model: %v", err)
	}

	return nil
}
