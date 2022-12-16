package aicmd

import (
	"encoding/json"
	"log"
	"os"
)

func LoadConfig() (config *Config, err error) {
	config, err = LoadConfigFromFile("config.json")
	if err != nil {
		log.Println("Failed to load config.json, using default config")
		log.Println("Creating new config.json with standard settings")
		config = NewDefaultConfig()
		err = SaveConfigToFile("config.json", config)
		if err != nil {
			return
		}
	}
	return
}

type Config struct {
	// The API key to use for the OpenAI API
	APIKey string `json:"api_key"`
	// The model to use for the OpenAI API
	Model string `json:"model_name"`
	// The temperature to use for the OpenAI API
	Temperature float64 `json:"temperature"`
	// The maximum number of tokens to use for the OpenAI API
	MaxTokens int `json:"max_tokens"`
	// The language model to use for the OpenAI API
}

func NewDefaultConfig() *Config {
	return &Config{
		APIKey:      "",
		Model:       "text-davinci-003",
		Temperature: 0.8,
		MaxTokens:   256,
	}
}

func LoadConfigFromFile(filename string) (*Config, error) {

	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	configFromJson := Config{}
	err = json.NewDecoder(configFile).Decode(&configFromJson)
	if err != nil {
		return nil, err
	}
	return &configFromJson, nil
}

func SaveConfigToFile(filename string, config *Config) error {
	configFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer configFile.Close()

	err = json.NewEncoder(configFile).Encode(config)
	if err != nil {
		return err
	}
	return nil
}
