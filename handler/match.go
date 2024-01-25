package handler

import (
	"joust/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type MatchHandler struct {
	MatchService *service.MatchService
}

func (h *MatchHandler) Create(c echo.Context) error {
	homeID, err := strconv.Atoi(c.FormValue("home"))
	if err != nil {
		return err
	}

	awayID, err := strconv.Atoi(c.FormValue("away"))
	if err != nil {
		return err
	}

	match, err := h.MatchService.SimMatch(c.Request().Context(), uint(homeID), uint(awayID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, match)
}
