package web_controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func FragmentPage(c echo.Context) error {
	return c.Render(http.StatusOK, "fragment", nil)
}
