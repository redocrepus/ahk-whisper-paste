package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Config struct {
	OpenapiKey     string
	AutoHotKeyExec string
}

// This function writes a given text to the clipboard using Windows' `clip` command
func writeTextToClipboard(text string) error {
	cmd := exec.Command("cmd", "/c", "echo|set /p="+text+"| clip")
	err := cmd.Run()
	return err
}

func main() {

	// Open a file for logging
	logFile, e := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(e)
	}
	defer logFile.Close()

	// Set the log output to the file
	log.SetOutput(logFile)
	log.Println("")
	log.Println("")
	log.Println("========================================")
	log.Println("Starting whisper-autohotkey")
	err := assertThatConfigFileExists()
	if err != nil {
		log.Fatal("Error when creating config file: ", err)
	}

	content, err := readConfigFile()
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during JSON parse: ", err)
	}

	if config.OpenapiKey == "" {
		log.Fatal("Please provide your OpenAI key in the file config.json")
	}

	// argLength := len(os.Args[1:])
	inputFileName := "rec.mp3"
	// if argLength > 1 {
	// 	inputFileName = os.Args[1:][1]
	// 	log.Println("Processing file " + inputFileName)
	// 	stats, err := os.Stat(inputFileName)
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		log.Fatal("Input file does not exist")
	// 	} else {
	// 		log.Printf("File size %v", stats.Size())
	// 	}
	// }

	text, err := Transcribe(inputFileName, config)
	if err != nil {
		log.Fatal("Cannot transcribe text: ", err)
		return
	}

	log.Println("Prompt:\n  " + text)

	err = writeTextToClipboard(text)
	if err != nil {
		log.Fatal("Failed to write text to clipboard:", err)
	}

	log.Println("Text copied to clipboard")

	ahkScript := `Send, ^v ; Ctrl+V for paste
ExitApp ; Exit after executing
`

	// Assuming you have AutoHotKey installed and `paste.ahk` is in the same directory.
	_, err = RunCommand(config, ahkScript)
	if err != nil {
		log.Fatal("Cannot run AutoHotKey command", err)
	}

	/*
		log.Println("Prompt:\n  " + text)
		command, err := BuildCommand(config, text)
		if err != nil {
			log.Fatal("Cannot interpret prompt", err)
			return
		}

		fmt.Println("Running:\n  " + command)
		output, err := RunCommand(config, command)
		if err != nil {
			log.Fatal("Cannot run command", err)
		}
		log.Println("Output:\n  " + output)
	*/
}

func readConfigFile() ([]byte, error) {
	content, err := os.ReadFile("./config.json")
	return content, err
}

func assertThatConfigFileExists() error {
	if !exists("./config.json") {
		template, err := os.ReadFile("./config.template.json")
		if err != nil {
			return fmt.Errorf("cannot read template config file: %w", err)
		}
		err = os.WriteFile("./config.json", template, 0666)
		if err != nil {
			return fmt.Errorf("cannot write new config file: %w", err)
		}
		return nil
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
