package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(inputFileName string, config Config) (string, error) {

	// Define the prompt variable and the language symbol
	prompt := "This is a transcription in English, mainly about programming, coding and software development."
	promptFileName := "transcriptionPrompt.txt"
	languageSymbol := "en" // default language symbol

	argLength := len(os.Args[1:])
	if argLength >= 2 { // Check if at least two arguments are provided
		languageSymbol = os.Args[1]
		promptFileName = os.Args[2]

		// log.Println("Processing file " + promptFileName)
		stats, err := os.Stat(promptFileName)
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("Prompt file does not exist")
		} else {
			log.Printf("File size %v", stats.Size())
		}
		log.Printf("Language Symbol: %s", languageSymbol)
	} else {
		log.Println("Insufficient arguments. Using default values.")
	}

	content, err := os.ReadFile(promptFileName)
	if err != nil {
		log.Fatal("Failed reading file: ", err)
	}
	prompt = string(content)

	c := openai.NewClient(config.OpenapiKey)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		Prompt:   prompt,
		Language: languageSymbol,
		FilePath: inputFileName,
	}
	response, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}

	return response.Text, nil
}
