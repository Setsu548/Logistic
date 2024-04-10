package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Setsu548/Logistic/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to load .env file")
	}

	db, err := connectDb()
	if err != nil {
		log.Fatal("Error to connect to database: ", err)
	}
	defer db.Close()

	e := echo.New()
	routes.InitRoutes(e, db)
	e.Start(":" + os.Getenv("HOST_PORT"))
}

func connectDb() (*gorm.DB, error) {
	// Obtener las variables de entorno para la conexión a la base de datos
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construir la cadena de conexión
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos MySQL
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
