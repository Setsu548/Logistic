package controllers

import (
	"net/http"

	"github.com/Setsu548/Logistic/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProducController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProducController {
	return &ProducController{DB: db}
}

func (cc *ProducController) CreateProduct(c echo.Context) error {
	product := new(models.Product)

	err := c.Bind(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error to code product")
	}

	var discount float32
	if product.Delivery.ProductAmount > 10 {
		switch product.Type {
		case "truck":
			discount = 0.05
			product.Delivery.Shipping_price -= product.Delivery.Shipping_price * discount
		case "maritime":
			discount = 0.03
			product.Delivery.Shipping_price -= product.Delivery.Shipping_price * discount
		default:
			return c.JSON(http.StatusBadRequest, "type of transport not correct")
		}
	}

	db := cc.DB.Create(&product)
	if db.Error != nil {
		return c.JSON(http.StatusBadRequest, db.Error.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

func (cc *ProducController) GetProducts(c echo.Context) error {
	var products []models.Product
	cc.DB.Find(&products)
	return c.JSON(http.StatusOK, products)
}
