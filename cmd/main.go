// main.go - application entry point
package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/DucTran999/dbkit"
	"github.com/DucTran999/dbkit/config"
	"github.com/DucTran999/go-clean-archx/internal/controller"
	"github.com/DucTran999/go-clean-archx/internal/repository"
	"github.com/DucTran999/go-clean-archx/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Setup DB
	conn, err := setupDB()
	if err != nil {
		log.Fatalln("setup db err:", err)
	}

	// Dependency Injection (DI): repo → usecase → controller
	productRepo := repository.NewProductRepository(conn.DB())
	productUC := usecase.NewProductUsecase(productRepo)
	productCtrl := controller.NewProductController(productUC)

	// Init router
	r := gin.Default()
	r.POST("/products", productCtrl.CreateProduct)

	// Start server
	addr := net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT"))
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func setupDB() (dbkit.Connection, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	conn, err := dbkit.NewPostgreSQLConnection(config.PostgreSQLConfig{
		Config: config.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			TimeZone: os.Getenv("DB_TIMEZONE"),
		},
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
