package route

import (
	"echo-gorm/api"
	"echo-gorm/controllers"
	"echo-gorm/middlewares"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", api.Home)

	e.GET("/users", api.GetUsers, middlewares.IsAuthenticated)

	e.GET("users/:username", api.GetUser, middlewares.IsAuthenticated)
	e.POST("users/:username", api.Authenticate, middlewares.IsAuthenticated)
	e.POST("users", api.CreateUser, middlewares.IsAuthenticated)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/product/", api.GetProducts)
	e.POST("/product/", api.CreateProduct)
	e.GET("/product/:idProduct", api.GetProduct)
	e.PUT("/product/:idProduct", api.UpdateProduct)
	e.DELETE("/product/:idProduct", api.DeleteProduct)

	return e
}
