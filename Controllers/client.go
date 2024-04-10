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

	dx := cc.DB.Create(&client)
	if dx.Error != nil {
		return c.JSON(http.StatusBadRequest, dx.Error.Error())
	}

	return c.JSON(http.StatusCreated, client)
}

func (cc *ClientController) GetClients(c echo.Context) error {
	var clients []models.Client
	cc.DB.Find(&clients)

	return c.JSON(http.StatusOK, clients)
}
