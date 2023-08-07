package models

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
