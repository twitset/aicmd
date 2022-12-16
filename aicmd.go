package aicmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func QueryGPT3(prompt string, config *Config) (answer string, err error) {

	// create a string that appends text
	jsonBody := map[string]interface{}{
		"model":       config.Model,
		"prompt":      prompt,
		"temperature": config.Temperature,
		"max_tokens":  config.MaxTokens,
	}

	jsonData, err := json.Marshal(jsonBody)

	// set up the request
	url := "https://api.openai.com/v1/completions"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIKey)

	// make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	answerBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	textCompletion := TextCompletion{}
	err = json.Unmarshal(answerBytes, &textCompletion)

	if err != nil {
		return
	}

	answer = textCompletion.Choices[0].Text
	answerWithoutNewlines := strings.Replace(answer, "\\n\\n", "", 1)

	return answerWithoutNewlines, nil
}

type TextCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
