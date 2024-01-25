package handler

import (
	"joust/data"
	"joust/view"

	"github.com/labstack/echo"
)

type KnightHandler struct {
	KnightRepo *data.KnightRepo
}

func (h *KnightHandler) ReadById(c echo.Context) error {
	println("handler")
	ctx := c.Request().Context()

	id, err := GetValidId(c)
	if err != nil {
		return err
	}

	knight, err := h.KnightRepo.ReadById(ctx, id)
	if err != nil {
		return err
	}

	return view.KnightPage(knight).Render(ctx, c.Response().Writer)
}
