package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OpenAIResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []OpenAIChoice `json:"choices"`
}

type OpenAIChoice struct {
	Index        int           `json:"index"`
	Message      OpenAIMessage `json:"message"`
	LogProbs     any           `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type OpenAIMessage struct {
	Role        string `json:"role"`
	Content     string `json:"content"`
	Refusal     any    `json:"refusal"`
	Annotations []any  `json:"annotations"`
}

func openai(commitMessage string, systemPrompt string, openaiApiKey string) (string, error) {
	payload := map[string]any{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{
				"role":    "developer",
				"content": systemPrompt,
			},
			{
				"role":    "user",
				"content": commitMessage,
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "Error: marshalling JSON", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "Error: creating request", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openaiApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error: making request", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error: reading response", err
	}

	var response OpenAIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "Error: parsing JSON", err
	}

	if len(response.Choices) == 0 {
		return "Error: No messages returned from LLM", errors.New("No messages")
	}

	return response.Choices[0].Message.Content, nil
}
