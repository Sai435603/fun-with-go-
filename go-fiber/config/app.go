package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber/models"
)

var app *fiber.App

func GetApp() *fiber.App {
	app = fiber.New()
	return app
}

var db []models.User

func GetDbInstance() *[]models.User {
	return &db
}
