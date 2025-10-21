package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/conneroisu/groq-go"
	"github.com/fatih/color"
)

var models []string = getModels()

var boldGreen *color.Color = color.New(color.FgGreen, color.Bold)

func setup() {
	// setup will
	// 1. ask for api key
	// 2. let the user choose model
	// dump the config to ~/.config/.fast-ai
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Welcome! Please configure settings for Fast AI.")
	var apiKey string
	boldGreen.Print("Enter your Groq API Key: ")
	apiKey, _ = reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	// int or string
	var modelChoice int
	boldGreen.Println("Choose a model:")
	for i, model := range models {
		boldGreen.Printf("%d. %s\n", i+1, model)
	}
	boldGreen.Print("Enter choice (1-", len(models), "): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	modelChoice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Invalid choice, please enter a number.")
		return
	}

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

	// if one argument is passed, we'll use it as the prompt
	if len(os.Args) == 2 {
		prompt := os.Args[1]

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
	// some global commands
	// exit or e - to exit
	// config to re-run setup

	var messages []groq.ChatCompletionMessage

	fmt.Println()
	fmt.Println("Type " + color.YellowString("'exit'") + " or " + color.YellowString("'e'") + " to exit, " + color.YellowString("'config'") + " to re-run setup")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		var userInput string

		fmt.Print(color.MagentaString("You: "))

		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" || userInput == "e" {
			fmt.Println("Exiting...")
			break
		}
		if userInput == "config" {
			setup()
			continue
		}

		if userInput == "" {
			continue
		}

		messages = append(messages, groq.ChatCompletionMessage{
			Role:    groq.RoleUser,
			Content: userInput,
		})
		response, err := generateResponse(messages)
		if err != nil {
			fmt.Println("Error generating response:", err)
			return
		}

		boldGreen.Print("AI: ")
		fmt.Println(" " + response)
		messages = append(messages, groq.ChatCompletionMessage{
			Role:    groq.RoleAssistant,
			Content: response,
		})
		fmt.Println()
	}

}
