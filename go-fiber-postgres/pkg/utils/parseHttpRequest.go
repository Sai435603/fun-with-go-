package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber-postgres/pkg/models"
)

func ParseData(stock *models.Stock, c *fiber.Ctx) error {
	return c.BodyParser(stock)
}
