package search

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func (svc *service) Handler(c echo.Context) error {
	cards, err := svc.guruRepo.ListCards()

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

	return c.JSON(http.StatusOK, cards)
}
