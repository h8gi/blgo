package api_controllers

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
	name := c.FormValue("name")
	contents := c.FormValue("contents")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty name")
	}
	if contents == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty contents")
	}

	f := &models.Fragment{
		Name:     name,
		Contents: contents,
	}
	if err := DB.Create(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, f)
}

func GetFragment(c echo.Context) error {
	name := c.Param("name")
	f := new(models.Fragment)
	if DB.First(f, "name = ?", name).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, f)
}

func UpdateFragment(c echo.Context) error {
	name := c.Param("name")
	contents := c.FormValue("contents")
	f := new(models.Fragment)
	if DB.First(f, "name = ?", name).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := DB.Model(f).Update("contents", contents).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}

func DeleteFragment(c echo.Context) error {
	name := c.Param("name")
	f := new(models.Fragment)
	if DB.First(f, "name = ?", name).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err := DB.Delete(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}
