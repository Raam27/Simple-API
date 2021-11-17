package controllers

import (
	"net/http"

	"echo-gorm/helpers"

	"github.com/labstack/echo"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := helpers.CheckLogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "berhasil login")
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}
