package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Type     string `json:"type"`
	Delivery Delivery
}
