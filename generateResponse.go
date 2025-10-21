package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/conneroisu/groq-go"
	"github.com/joho/godotenv"
)

func generateResponse(model string, prompt string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

	}
	api := os.Getenv("groq_API")
	client, err := groq.NewClient(api)
	if err != nil {
		fmt.Println("Error creating Groq client:", err)
		return "", err
	}

	response, err := client.ChatCompletion(
		context.Background(),
		groq.ChatCompletionRequest{
			Model: groq.ChatModel(model), // Select your model
			Messages: []groq.ChatCompletionMessage{
				{Role: groq.RoleSystem,
					Content: "You'll always respond in a concise to the point manner, answer what is being asked and nothing more, avoid using markdown formatting."},
				{
					Role:    groq.RoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	// remove everything beween <think> and </think> tags from the response, remove the tags and content betweent it
	responseText := response.Choices[0].Message.Content

	startTag := "<think>"
	endTag := "</think>"
	for {
		startIndex := strings.Index(responseText, startTag)
		endIndex := strings.Index(responseText, endTag)
		if startIndex == -1 || endIndex == -1 || endIndex < startIndex {
			break
		}
		responseText = responseText[:startIndex] + responseText[endIndex+len(endTag):]
	}

	return strings.TrimSpace(responseText), nil

}
