package middlewares

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(Cek()),
})

func Cek() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// fmt.Println(os.Getenv("secret"))
	a := os.Getenv("secret")
	return a
}
