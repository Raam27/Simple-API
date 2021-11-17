package api

import (
	"echo-gorm/db"
	"echo-gorm/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []model.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := db.DbManager()
	username := c.Param("username")
	user := model.User{}
	db.First(&user, "username = ?", username)
	if user.Username == "" {
		return c.NoContent(http.StatusNotFound)
	}
	usernameDB := user.Username
	firstName := "Coba"
	email := fmt.Sprintf("%s@localhost.com", usernameDB)
	keycloak := model.NewKeycloak(usernameDB, firstName, email)
	return c.JSON(http.StatusOK, keycloak)
}

func Authenticate(c echo.Context) error {
	var pass model.Password
	db := db.DbManager()
	if c.Request().Body == nil {
		c.String(http.StatusUnauthorized, "body kosong")
	}
	err := json.NewDecoder(c.Request().Body).Decode(&pass)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}
	username := c.Param("username")
	user := model.User{}
	db.First(&user, "username = ?", username)
	// fmt.Println(pass.Password)
	// fmt.Println(user.Password)
	if pass.Password == user.Password {
		return c.NoContent(http.StatusOK)
	}
	return c.String(http.StatusUnauthorized, "password salah")

}

func CreateUser(c echo.Context) error {
	user := model.User{}
	db := db.DbManager()
	if c.Request().Body == nil {
		c.String(http.StatusUnauthorized, "body kosong")
	}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}
	validate := validator.New()
	err = validate.Struct(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Username = strings.ToLower(user.Username)
	fmt.Println(user.Username)
	db.Create(&user)
	return c.JSON(http.StatusCreated, "berhasil membuat")
}
