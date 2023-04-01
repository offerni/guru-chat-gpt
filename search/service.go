package search

import "github.com/offerni/guruchatgpt"

type Service struct {
	Credentials guruchatgpt.Credentials
	GuruRepo    guruchatgpt.GuruRepository
	OpenAIRepo  guruchatgpt.OpenAIRepository
}

type NewServiceOpts struct {
	Credentials      guruchatgpt.Credentials
	GuruRepository   guruchatgpt.GuruRepository
	OpenAiRepository guruchatgpt.OpenAIRepository
}

func NewService(opts NewServiceOpts) (*Service, error) {
	return &Service{
		Credentials: opts.Credentials,
		GuruRepo:    opts.GuruRepository,
		OpenAIRepo:  opts.OpenAiRepository,
	}, nil
}
