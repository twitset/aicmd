package main

import (
	"aicmd"
	"flag"
	"log"
	"os"
)

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	var tokens int
	var temperature float64
	var model string

	flag.IntVar(&tokens, "max_tokens", 256, "The maximum number of tokens to generate *optional")
	flag.Float64Var(&temperature, "temperature", 0.9999, "The temperature to use for the AI *optional")
	flag.StringVar(&model, "model", "text-davinci-003", "The model to use for the AI *optional")
	flag.Parse()

	prompt := flag.Arg(0) // the last argument on the command line is the prompt

	config, err := aicmd.LoadConfig()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	if IsFlagPassed("tokens") {
		config.MaxTokens = tokens
	}
	if IsFlagPassed("temperature") {
		config.Temperature = temperature
	}
	if IsFlagPassed("model") {
		config.Model = model
	}

	if config.APIKey == "" {
		os.Stderr.WriteString("API key is missing. Set it in config.json")
		os.Exit(1)
	}

	if prompt == "" {
		os.Stdout.WriteString("Write a task for the AI.")
		os.Exit(1)
	}

	answer, err := aicmd.QueryGPT3(prompt, config)

	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.WriteString(answer)
	if err != nil {
		log.Fatal(err)
	}

}
