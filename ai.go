package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func getModels() []string {
	openai120B := "openai/gpt-oss-120b"   // 500 T/s
	openai20B := "openai/gpt-oss-20b"     // 1000 T/s
	llama70B := "llama-3.3-70b-versatile" // 280 T/s
	compound := "groq/compound"           // 450 T/s
	qwen := "qwen/qwen3-32b"              // 400 T/s

	models := []string{
		openai120B,
		openai20B,
		llama70B,
		compound,
		qwen,
	}
	return models
}

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

	configContent := fmt.Sprintf("API_KEY=%s\nMODEL=%s\n", apiKey, selectedModel)

	configPath := os.Getenv("HOME") + "/.config/.fast-ai"
	err := os.WriteFile(configPath, []byte(configContent), 0600)

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

	prompt := os.Args[1]

	response, err := generateResponse(models[2], prompt)
	if err != nil {
		fmt.Println("Error generating response:", err)
		return
	}

	boldGreen.Print("Response:")
	fmt.Println(" " + response)

}
