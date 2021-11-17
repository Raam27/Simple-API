package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" gorm:"primaryKey" validate:"required"`
	Password string `json:"password" validate:"required"`
}
