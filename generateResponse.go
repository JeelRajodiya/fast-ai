package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/conneroisu/groq-go"
)

func generateResponse(prompt interface{}) (string, error) {
	config, err := getConfig()

	if err != nil {
		return "", fmt.Errorf("error getting config: %w", err)
	}

	api := config.GROQ_API_KEY
	model := config.Model

	client, err := groq.NewClient(api)
	if err != nil {
		fmt.Println("Error creating Groq client:", err)
		return "", err
	}

	var messages []groq.ChatCompletionMessage
	systemPrompt := "You'll always respond in a concise to the point manner, answer what is being asked and nothing more, never use markdown in your responses."
	switch p := prompt.(type) {
	case string:
		messages = []groq.ChatCompletionMessage{
			{Role: groq.RoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    groq.RoleUser,
				Content: p,
			},
		}
	case []groq.ChatCompletionMessage:
		// append system prompt at the start
		messages = append([]groq.ChatCompletionMessage{
			{
				Role:    groq.RoleSystem,
				Content: systemPrompt,
			},
		}, p...)

	default:
		return "", fmt.Errorf("invalid prompt type: %T", p)
	}

	response, err := client.ChatCompletion(
		context.Background(),
		groq.ChatCompletionRequest{
			Model:    groq.ChatModel(model), // Select your model
			Messages: messages,
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
