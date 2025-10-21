package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/conneroisu/groq-go"
	"github.com/fatih/color"
)

var boldGreen *color.Color = color.New(color.FgGreen, color.Bold)

func main() {

	// check if file ~/.config/.fast-ai exists, if it does not exist we'll run the setup
	config, err := getConfig()

	if os.IsNotExist(err) || (len(os.Args) == 2 && os.Args[1] == "--config") {
		setup()
		fmt.Println("Setup complete! You can now use the ai command.")
		return
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

	fmt.Println("Config loaded! using Model:", color.YellowString(config.Model))
	fmt.Println("Type " + color.YellowString("'exit'") + " or " + color.YellowString("'e'") + " to exit, " + color.YellowString("'config'") + " to re-run setup, or use " + color.YellowString("ai --config") + " to re-run the setup")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		var userInput string

		fmt.Print(color.MagentaString("You: "))

		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" || userInput == "e" || userInput == "quit" || userInput == "q" {
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
