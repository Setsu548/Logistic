package models

import "github.com/jinzhu/gorm"

type Delivery struct {
	gorm.Model
	ProductAmount  int     `json:"product_amount"`
	Register_date  string  `json:"register_date"`
	Delivery_date  string  `json:"delivery_date"`
	Warehouse      string  `json:"warehouse"`
	Shipping_price float32 `json:"shipping_price"`
	Document       string  `json:"document"`
	Guide_number   string  `json:"guide_number"`
	ClientId       int     `json:"client_id"`
	ProductId      uint    `json:"product_id"`
}
