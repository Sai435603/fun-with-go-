package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	r.SetUpRoutes(app)

	app.Listen(":8000")

}
