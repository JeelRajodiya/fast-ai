package main

import (
	"encoding/json"
	"fmt"
	"os"
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
