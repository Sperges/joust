package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var ErrNoName = errors.New("no name specified")
var ErrNoPassword = errors.New("no password specified")

func GetValidId(c echo.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}

	if id < 1 {
		return 0, c.String(http.StatusBadRequest, "knight id cannot be less than 1")
	}

	return uint(id), nil
}
