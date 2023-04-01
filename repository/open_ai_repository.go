package repository

type OpenAIRepo struct {
	Chat        Chat // should be its own repo but it's late and I'm lazy
	Credentials Credentials
}

type NewOpenAIRepoOpts struct {
	Chat        Chat
	Credentials Credentials
}

type Credentials struct {
	BearerToken string
}

type Chat struct {
	CompletionRequestOpts *OpenAIChatCompletionRequest
}

func NewOpenAIRepository(opts NewOpenAIRepoOpts) (*OpenAIRepo, error) {
	return &OpenAIRepo{
		Chat:        opts.Chat,
		Credentials: opts.Credentials,
	}, nil
}

type OpenAIChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []Message               `json:"messages"`
	Temperature      *float64                `json:"temperature,omitempty"`
	TopP             *float64                `json:"top_p,omitempty"`
	N                *int                    `json:"n,omitempty"`
	Stream           *bool                   `json:"stream,omitempty"`
	Stop             *interface{}            `json:"stop,omitempty"`
	MaxTokens        *int                    `json:"max_tokens,omitempty"`
	PresencePenalty  *float64                `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64                `json:"frequency_penalty,omitempty"`
	LogitBias        *map[string]interface{} `json:"logit_bias,omitempty"`
	User             *string                 `json:"user,omitempty"`
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
