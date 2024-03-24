package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	GPT3Dot5Turbo = "gpt-3.5-turbo"
	GPT4          = "gpt-4"
)

type GPT struct {
	ApiKey     string
	HttpClient *http.Client
}

func NewGPT(apiKey string) *GPT {
	return &GPT{
		ApiKey:     apiKey,
		HttpClient: &http.Client{},
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	Messages     []Message     `json:"messages"`
	Model        string        `json:"model"`
	Functions    []GPTFunction `json:"functions"`
	FunctionCall string        `json:"function_call"`
}

const apiURL = "https://api.openai.com/v1/chat/completions"

func (client GPT) NewChatCompletion(prompt string, model string, function ...GPTFunction) (response string, err error) {
	payload := Payload{
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Model:        model,
		Functions:    function,
		FunctionCall: "auto",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+client.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = string(body)
	return
}
