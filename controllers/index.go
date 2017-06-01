package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
