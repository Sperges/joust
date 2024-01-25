package handler

import (
	"joust/data"
	"joust/model"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	Repo data.Repo[model.User]
}

func (h UserHandler) Create(c echo.Context) error {
	name := c.FormValue("name")
	if name == "" {
		c.NoContent(http.StatusBadRequest)
		return ErrNoName
	}

	password := c.FormValue("password")
	if password == "" {
		c.NoContent(http.StatusBadRequest)
		return ErrNoPassword
	}

	user := model.User{
		Name:     name,
		Password: password,
	}

	if err := h.Repo.Create(c.Request().Context(), &user); err != nil {
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
