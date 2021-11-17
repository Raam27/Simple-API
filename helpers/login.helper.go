package helpers

import (
	"echo-gorm/db"
	"echo-gorm/model"
	"fmt"
)

func CheckLogin(username, password string) (bool, error) {
	db := db.DbManager()
	user := model.User{}
	db.First(&user, "username = ?", username)

	match, err := CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("hash and pass doesn't match")
		return false, err
	}
	return true, nil
}
