package repository

import (
	"encoding/json"
	"net/http"

	"github.com/offerni/guruchatgpt"
)

func (gr *GuruRepo) ListCards() (*guruchatgpt.ListGuruCards, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.getguru.com/api/v1/cards", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(gr.Username, gr.Password)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cards []guruchatgpt.Card
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		return nil, err
	}

	return &guruchatgpt.ListGuruCards{
		Cards: cards,
	}, nil
}
