package api

import (
	"echo-gorm/db"
	"echo-gorm/model"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func CreateProduct(c echo.Context) error {
	product := model.Product{}
	db := db.DbManager()
	if c.Request().Body == nil {
		c.String(http.StatusUnauthorized, "body kosong")
	}
	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}
	db.Create(&product)
	return c.JSON(http.StatusCreated, "berhasil membuat")
}

func UpdateProduct(c echo.Context) error {
	db := db.DbManager()
	id := c.Param("idProduct")
	product := model.Product{}
	db.First(&product, id)
	if c.Request().Body == nil {
		c.String(http.StatusUnauthorized, "body kosong")
	}
	err := json.NewDecoder(c.Request().Body).Decode(&product)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
	}
	db.Save(&product)
	return c.JSON(http.StatusOK, "berhasil update")
}

func DeleteProduct(c echo.Context) error {
	db := db.DbManager()
	id := c.Param("idProduct")
	product := model.Product{}
	db.Delete(&product, id)
	return c.JSON(http.StatusNoContent, "berhasil menghapus")
}

func GetProduct(c echo.Context) error {
	db := db.DbManager()
	id := c.Param(("idProduct"))
	product := model.Product{}
	db.First(&product, id)
	return c.JSON(http.StatusOK, product)
}

func GetProducts(c echo.Context) error {
	db := db.DbManager()
	products := []model.Product{}
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}
