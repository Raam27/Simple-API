package db

import (
	"echo-gorm/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_PORT, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	if err != nil {
		panic("DB Connection Error")
	}
	// defer db.Close()
	// db.AutoMigrate(&model.User{})

}

func DbManager() *gorm.DB {
	return db
}
