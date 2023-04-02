package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/offerni/guruchatgpt"
)

const openAIBaseUrl = "https://api.openai.com/v1"

func (or *OpenAIRepo) ChatCompletion(
	opts guruchatgpt.ChatCompletionRequestOpts,
) (*guruchatgpt.ChatCompletionResponse, error) {

	reqBody := OpenAIChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: opts.Message,
			},
		},
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, openAIBaseUrl+"/chat/completions", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+opts.Credentials.BearerToken)
	req.Header.Set("Content-Type", "application/json")

	log.Println("request body:", bytes.NewBuffer(reqBodyBytes))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var responseMessage guruchatgpt.ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	if err != nil {
		return nil, err
	}

	return &responseMessage, nil
}
