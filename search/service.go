package search

import "github.com/offerni/guruchatgpt"

type service struct {
	guruRepo guruchatgpt.GuruRepository
}

type NewServiceOpts struct {
	GuruRepository guruchatgpt.GuruRepository
}

func NewService(opts NewServiceOpts) (*service, error) {
	return &service{
		guruRepo: opts.GuruRepository,
	}, nil
}
