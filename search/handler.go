package search

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (svc *service) Handler(c echo.Context) error {
	cards, err := svc.guruRepo.ListCards()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, cards)
}
