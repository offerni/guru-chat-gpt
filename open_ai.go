package guruchatgpt

import "github.com/labstack/echo/v4"

type ChatCompletionRequestOpts struct {
	Message     string
	Credentials Credentials
	Dataset     any
}

type Credentials struct {
	BearerToken string
}

type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   ChatCompletionUsage    `json:"usage"`
}

type ChatCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRepository interface {
	ChatCompletion(echo.Context, ChatCompletionRequestOpts) error
}
