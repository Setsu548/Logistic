package routes

import (
	controllers "github.com/Setsu548/Logistic/Controllers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	clientController := controllers.NewClientController(db)
	e.POST("/client", clientController.CreateClient)
}
