package search

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/offerni/guruchatgpt"
)

func (svc *Service) Handler(c echo.Context) error {
	cards, err := svc.GuruRepo.ListCards()
	message := c.QueryParam("message")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	jsonCards, err := json.Marshal(cards)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	stringfiedCards := string(jsonCards)
	spew.Dump(stringfiedCards)

	chatGptResponse, err := svc.OpenAIRepo.ChatCompletion(guruchatgpt.ChatCompletionRequestOpts{
		Message: message,
		Credentials: guruchatgpt.Credentials{
			BearerToken: svc.Credentials.BearerToken,
		},
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, chatGptResponse)
}
