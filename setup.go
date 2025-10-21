package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"

	"os"
	"strconv"
	"strings"
)

func setup() {
	// setup will
	// 1. ask for api key
	// 2. let the user choose model
	// dump the config to ~/.config/.fast-ai
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Welcome! Please configure settings for Fast AI.")

	oldConfig, configErr := getConfig()
	if configErr == nil {
		fmt.Println("Current configuration found:")
		fmt.Println("GROQ API Key: " + oldConfig.GROQ_API_KEY)
		fmt.Println("Model: " + oldConfig.Model)
		fmt.Println()
	}

	models := getModels()
	var apiKey string
	var underline = color.New(color.Underline)
	fmt.Println("Get API key from: ", underline.Sprint("https://console.groq.com/keys"))
	if configErr == nil {
		boldGreen.Print("Enter your Groq API Key (press enter to keep current): ")
	} else {
		boldGreen.Print("Enter your Groq API Key: ")
	}

	apiKeyInput, _ := reader.ReadString('\n')
	apiKeyInput = strings.TrimSpace(apiKeyInput)
	if apiKeyInput == "" && configErr == nil {
		apiKey = oldConfig.GROQ_API_KEY
	} else {
		apiKey = apiKeyInput
	}
	fmt.Println()

	fmt.Println("Choose a model:")
	// find max length of model name
	maxLen := 0
	for _, model := range models {
		if len(model.Name) > maxLen {
			maxLen = len(model.Name)
		}
	}
	for i, modelInfo := range models {
		yellow := color.New(color.FgHiYellow).SprintFunc()
		modelName := boldGreen.Sprintf("%d. %s", i+1, modelInfo.Name)
		padding := maxLen - len(modelInfo.Name) + 5
		fmt.Println(modelName, strings.Repeat(" ", padding), yellow(strconv.Itoa(modelInfo.Speed)+" tokens/sec"))
	}
	fmt.Println()
	fmt.Print(color.MagentaString(fmt.Sprintf("Enter choice (1-%d), press enter to use the default: ", len(models))))
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	var selectedModel ModelInfo
	if choiceStr == "" {
		// find default model and set it
		for _, model := range models {
			if model.Name == DEFAULT_MODEL {
				selectedModel = model
				break
			}
		}
	} else {
		modelChoice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice, please enter a number.")
			return
		}

		if modelChoice < 1 || modelChoice > len(models) {
			fmt.Println("Invalid choice")
			return
		}

		selectedModel = models[modelChoice-1]
	}

	boldGreen.Println("Selected model:", selectedModel.Name)

	// write to config file
	config := Config{
		GROQ_API_KEY: apiKey,
		Model:        selectedModel.Code,
	}
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error creating config JSON:", err)
		return
	}

	configPath := os.Getenv("HOME") + "/.config/.fast-ai"
	err = os.WriteFile(configPath, configJSON, 0600)

	if err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

	fmt.Println("Setup complete! Config saved to", configPath)
}
