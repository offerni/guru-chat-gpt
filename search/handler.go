package search

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/offerni/guruchatgpt"
)

func (svc *Service) Handler(c echo.Context) error {
	cards, err := svc.GuruRepo.ListCards()
	message := c.QueryParam("message")
	sessionID := c.QueryParam("sessionID")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	err = svc.OpenAIRepo.ChatCompletion(c, guruchatgpt.ChatCompletionRequestOpts{
		Message:   message,
		SessionID: sessionID,
		Credentials: guruchatgpt.Credentials{
			BearerToken: svc.Credentials.BearerToken,
		},
		Dataset: cards,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
