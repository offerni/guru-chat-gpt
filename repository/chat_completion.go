package repository

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/offerni/guruchatgpt"
)

const openAIBaseUrl = "https://api.openai.com/v1"

// Tweak here as needed to save some request tokens, the lower the number the
// smaller is the dataset generated so more generic the answer will be
const messageLimitHardcap = 100

func (or *OpenAIRepo) ChatCompletion(
	c echo.Context,
	opts guruchatgpt.ChatCompletionRequestOpts,
) error {

	// since the value is any as it could be any give dataset, so here we convert
	// it to the known type
	cards, ok := opts.Dataset.(*guruchatgpt.ListGuruCards)
	if !ok {
		return errors.New("type assertion failed on: opts.Dataset.(*guruchatgpt.ListGuruCards)")
	}

	jsonCards, err := json.Marshal(cards.Cards)
	if err != nil {
		return err
	}

	// Unfortunately gpt-3 API limits the message size to 4097 tokens which is not
	// a lot, I've even tried breaking the messages into smaller chunks but it
	// seems this limit is per payload and not per message a possible solution
	// could be finding a way to store this data accurately and safely so the api
	// could tap into that
	// spew.Dump(string(jsonCards))
	cardsDatasetSlice := splitString(string(jsonCards[:messageLimitHardcap]), messageLimitHardcap)

	// this is a mess, all this messages logic should be moved to the handler or
	// to an intermediate business logic layer or something but good enough for now
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

	reqBody := OpenAIChatCompletionRequest{
		Stream:   true,
		Model:    or.Chat.CompletionRequestOpts.Model,
		Messages: messages,
	}

	reqBodyBytes, err := json.Marshal(&reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, openAIBaseUrl+"/chat/completions", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}

	setRequestHeaders(req, opts)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	if opts.Message != "" {
		err = listenAndForwardEventStream(c, resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func listenAndForwardEventStream(c echo.Context, resp *http.Response) error {
	reader := bufio.NewReader(resp.Body)

	setResponseHeaders(c.Response())

	// recursion reading the bytes until it errors out/returns skipping the loop
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("error reading response: %s", err.Error())
				return err
			}

			fmt.Fprintf(c.Response(), "event: close\ndata: Connection Closed\n\n")
			return nil
		}

		if len(line) > 0 {
			msg := fmt.Sprintf(string(line))
			fmt.Fprintf(c.Response(), msg)
		}
	}
}

// TODO: move it to a utils file
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

func setRequestHeaders(req *http.Request, opts guruchatgpt.ChatCompletionRequestOpts) {
	req.Header.Set("Authorization", "Bearer "+opts.Credentials.BearerToken)
	req.Header.Set("Content-Type", "application/json")
}

func setResponseHeaders(response *echo.Response) {
	response.Header().Set(echo.HeaderContentType, "text/event-stream")
	response.Header().Set("Cache-Control", "no-cache")
	response.Header().Set("Connection", "keep-alive")
}
