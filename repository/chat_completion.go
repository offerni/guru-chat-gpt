package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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
				Role:    "system",
				Content: "You're a helpful that searches and provides answers as concisely as possible and links when available given these Guru Cards:" + opts.Dataset,
			},
			{
				Role:    "system",
				Content: "Cite your refference as Guru Card",
			},
			{
				Role:    "user",
				Content: opts.Message,
			},
		},
	}

	spew.Dump(opts.Dataset)

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

	// log.Println("request body:", bytes.NewBuffer(reqBodyBytes))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		spew.Dump(string(body))
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
