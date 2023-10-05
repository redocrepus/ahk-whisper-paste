package main

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(inputFileName string, config Config) (string, error) {

	// Define the prompt variable.
	prompt := "This is a transcription in English, mainly about programming, coding and software development."

	content, err := os.ReadFile("transtriptionPrompt.txt")
	if err != nil {
		log.Fatal("Failed reading file: ", err)
	}
	prompt = string(content)

	c := openai.NewClient(config.OpenapiKey)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		Prompt:   prompt,
		FilePath: inputFileName,
	}
	response, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return response.Text, nil
}
