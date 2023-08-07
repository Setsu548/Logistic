package controllers

import (
	"net/http"

	"github.com/Setsu548/Logistic/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ClientController struct {
	DB *gorm.DB
}

func NewClientController(db *gorm.DB) *ClientController {
	return &ClientController{DB: db}
}

func (cc *ClientController) CreateClient(c echo.Context) error {
	client := new(models.Client)

	err := c.Bind(client)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error to code the Client")
	}

	cc.DB.Create(client)

	return c.JSON(http.StatusCreated, client)
}
