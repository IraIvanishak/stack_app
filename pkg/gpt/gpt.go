package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// TODO expand
const (
	GPT3Dot5Turbo = "gpt-3.5-turbo"
	GPT4          = "gpt-4"
)

type GPT struct {
	ApiKey     string
	HttpClient *http.Client
	Model      string
}

func NewGPT(apiKey string, model string) *GPT {
	return &GPT{
		ApiKey:     apiKey,
		HttpClient: &http.Client{},
		Model:      model,
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

type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

type Choice struct {
	Index        int           `json:"index"`
	Message      ChoiceMessage `json:"message"`
	Logprobs     interface{}   `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type ChoiceMessage struct {
	Role         string       `json:"role"`
	Content      string       `json:"content"`
	FunctionCall FunctionCall `json:"function_call"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

const apiURL = "https://api.openai.com/v1/chat/completions"

func (client GPT) NewChatCompletion(prompt string, object interface{}, function ...GPTFunction) (err error) {
	payload := Payload{
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Model:        client.Model,
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

	// TODO: parse error response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response.Choices[0].Message.FunctionCall.Arguments), object)

	return
}
