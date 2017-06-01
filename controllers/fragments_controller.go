package controllers

import (
	"net/http"

	"github.com/h8gi/blgo/models"
	"github.com/labstack/echo"
)

func GetFragmentsList(c echo.Context) error {
	var fs []*models.Fragment
	DB.Find(&fs)
	return c.JSON(http.StatusOK, fs)
}

// TODO: Check form value. Don't use `c.bind`.
func PostFragment(c echo.Context) error {
	text := c.FormValue("text")
	f := &models.Fragment{
		Text: text,
	}
	if err := DB.Create(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, f)
}

func GetFragment(c echo.Context) error {
	id := c.Param("id")
	f := new(models.Fragment)
	if DB.First(f, "id = ?", id).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, f)
}

func UpdateFragment(c echo.Context) error {
	id := c.Param("id")
	text := c.Param("text")
	f := new(models.Fragment)
	if DB.First(f, "id = ?", id).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := DB.Model(f).Update("text", text).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}

func DeleteFragment(c echo.Context) error {
	id := c.Param("id")
	f := new(models.Fragment)
	if DB.First(f, "id = ?", id).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := DB.Delete(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}
