package routes

import (
	controllers "github.com/Setsu548/Logistic/Controllers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {

	clientController := controllers.NewClientController(db)
	client := e.Group("/client")
	client.POST("/", clientController.CreateClient)
	client.GET("/", clientController.GetClients)

	productController := controllers.NewProductController(db)
	p := e.Group("/product")
	p.POST("/", productController.CreateProduct)
	p.GET("/", productController.GetProducts)

}
