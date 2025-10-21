package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var models []string = getModels()

var boldGreen *color.Color = color.New(color.FgGreen, color.Bold)

func setup() {
	// setup will
	// 1. ask for api key
	// 2. let the user choose model
	// dump the config to ~/.config/.fast-ai

	fmt.Println()
	fmt.Println("Welcome! Please configure settings for Fast AI.")
	var apiKey string
	boldGreen.Print("Enter your Groq API Key: ")
	fmt.Scanln(&apiKey)

	// int or string
	var modelChoice int
	boldGreen.Println("Choose a model:")
	for i, model := range models {
		boldGreen.Printf("%d. %s\n", i+1, model)
	}
	boldGreen.Print("Enter choice (1-", len(models), "): ")
	fmt.Scanln(&modelChoice)

	if modelChoice < 1 || modelChoice > len(models) {
		fmt.Println("Invalid choice")
		return
	}

	selectedModel := models[modelChoice-1]
	boldGreen.Println("Selected model:", selectedModel)

	// write to config file

	config := Config{
		GROQ_API_KEY: apiKey,
		Model:        selectedModel,
	}
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error creating config JSON:", err)
		return
	}
	configContent := string(configJSON)

	configPath := os.Getenv("HOME") + "/.config/.fast-ai"
	err = os.WriteFile(configPath, []byte(configContent), 0600)

	if err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

	fmt.Println("Setup complete! Config saved to", configPath)

}

func main() {

	// check if file ~/.config/.fast-ai exists, if it does not exist we'll run the setup
	_, err := os.Stat(os.Getenv("HOME") + "/.config/.fast-ai")

	if os.IsNotExist(err) {
		setup()
	}

	// some global shortcuts
	// exit or e - to exit
	// config to re-run setup

	// if one argument is passed, we'll use it as the prompt
	if len(os.Args) == 2 {
		prompt := os.Args[1]
		fmt.Print(prompt)

		response, err := generateResponse(prompt)
		if err != nil {
			fmt.Println("Error generating response:", err)
			return
		}

		boldGreen.Print("Response:")
		fmt.Println(" " + response)
		return

	}

	// interactive mode

}
