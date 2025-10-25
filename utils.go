package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// model type
type ModelInfo struct {
	Name  string
	Code  string
	Speed int // tokens per second
}

var DEFAULT_MODEL = "Compound"

func getModels() []ModelInfo {
	models := []ModelInfo{
		{Name: "Qwen 3 32B", Code: "qwen/qwen3-32b", Speed: 400},
		{Name: "OpenAI GPT OSS 120B", Code: "openai/gpt-oss-120b", Speed: 500},
		{Name: "OpenAI GPT OSS 20B", Code: "openai/gpt-oss-20b", Speed: 1000},
		{Name: "Llama 3.3 70B", Code: "llama-3.3-70b-versatile", Speed: 280},
		{Name: "Llama 3.1 8B", Code: "llama-3.1-8b-instant", Speed: 560},
		{Name: "Compound", Code: "groq/compound", Speed: 450},
		{Name: "Compound Mini", Code: "groq/compound-mini", Speed: 450},
		{Name: "Llama 4 Maverick", Code: "meta-llama/llama-4-maverick-17b-128e", Speed: 600},
		{Name: "Llama 4 Scout", Code: "meta-llama/llama-4-scout-17b-16e-instruct", Speed: 750},
	}
	return models
}

type Config struct {
	GROQ_API_KEY string
	Model        string
}

func getConfig() (Config, error) {
	configPath := os.Getenv("HOME") + "/.config/.fast-ai"
	content, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	return config, nil
}
