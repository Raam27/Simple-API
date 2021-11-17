package model

import (
	"time"
)

type Product struct {
	Idnya     string    `gorm:"primary_key;column:idnya;type:int(32) NOT NULL AUTO_INCREMENT" json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price,string"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time
}
