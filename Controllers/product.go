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
