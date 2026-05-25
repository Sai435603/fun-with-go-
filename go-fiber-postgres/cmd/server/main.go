package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber-postgres/pkg/config"
	routes "github.com/sai435603/go-fiber-postgres/pkg/stocks-routes"
)

func main() {
	log.Println("Welcome to stocks app")
	config.ConnectDB()
	app := fiber.New()
	routes.RegisterStockRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
