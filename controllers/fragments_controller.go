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
	f := new(models.Fragment)
	if err := c.Bind(f); err != nil {
		return err
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
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, f)
}

func UpdateFragment(c echo.Context) error {
	f := new(models.Fragment)
	if err := c.Bind(f); err != nil {
		return err
	}
	if err := DB.Update(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}

func DeleteFragment(c echo.Context) error {
	id := c.Param("id")
	f := new(models.Fragment)
	if DB.First(f, "id = ?", id).RecordNotFound() {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
	}
	if err := DB.Delete(f).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, f)
}
