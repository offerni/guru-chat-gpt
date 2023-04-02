package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/offerni/guruchatgpt"
)

const openAIBaseUrl = "https://api.openai.com/v1"

// Tweak here as needed to save some request tokens, the lower the number the
// smaller is the dataset generated so more generic the answer will be
const messageLimitHardcap = 100

func (or *OpenAIRepo) ChatCompletion(
	opts guruchatgpt.ChatCompletionRequestOpts,
) (*guruchatgpt.ChatCompletionResponse, error) {

	// since the value is any as it could be any give dataset, so here we convert
	// it to the known type
	cards, ok := opts.Dataset.(*guruchatgpt.ListGuruCards)

	if !ok {
		return nil, errors.New("type assertion failed on: opts.Dataset.(*guruchatgpt.ListGuruCards)")
	}

	jsonCards, err := json.Marshal(cards.Cards)
	if err != nil {
		return nil, err
	}

	// Unfortunately gpt-3 API limits the message size to 4097 tokens which is not
	// a lot, I've even tried breaking the messages into smaller chunks but it
	// seems this limit is per payload and not per message a possible solution
	// could be finding a way to store this data accurately and safely so the api
	// could tap into that
	// spew.Dump(string(jsonCards))
	cardsDatasetSlice := splitString(string(jsonCards[:messageLimitHardcap]), messageLimitHardcap)

	messages := []Message{
		{
			Role:    "system",
			Content: "You're a helpful assistant that searches and provides answers as concisely as possible the given dataset known as Guru Cards",
		},
	}

	for _, cardString := range cardsDatasetSlice {
		messages = append(messages, Message{
			Role:    "system",
			Content: cardString,
		})
	}

	messages = append(messages, Message{
		Role:    "user",
		Content: opts.Message,
	})

	// spew.Dump(messages)

	reqBody := OpenAIChatCompletionRequest{
		Model:    or.Chat.CompletionRequestOpts.Model,
		Messages: messages,
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

func splitString(input string, size int) []string {
	var chunks []string
	for len(input) > 0 {
		if len(input) < size {
			size = len(input)
		}
		chunks = append(chunks, input[:size])
		input = input[size:]
	}
	return chunks
}
