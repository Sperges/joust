package handler

import (
	"joust/service"
	"net/http"

	"github.com/labstack/echo"
)

type GenerateHandler struct {
	GenerateService *service.GenerateService
}

func (h *GenerateHandler) Knights(c echo.Context) error {
	if err := h.GenerateService.Knights(c.Request().Context(), 100); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
