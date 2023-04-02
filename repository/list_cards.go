package repository

import (
	"encoding/json"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/offerni/guruchatgpt"
)

const guruApiBareUrl = "https://api.getguru.com/api/v1"

func (gr *GuruRepo) ListCards() (*guruchatgpt.ListGuruCards, error) {
	req, err := http.NewRequest(http.MethodGet, guruApiBareUrl+"/cards", nil)
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

	// Adding strict policy to remove any html tags or formatting since it doesn't
	// matter and it's just noise to feed to the Chat prompt
	p := bluemonday.StrictPolicy()

	for i, card := range cards {
		cards[i].Content = p.Sanitize(card.Content)
	}

	return &guruchatgpt.ListGuruCards{
		Cards: cards,
	}, nil
}
