package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber-postgres/pkg/config"
	"github.com/sai435603/go-fiber-postgres/pkg/models"
	"github.com/sai435603/go-fiber-postgres/pkg/utils"
)

func GetStocks(c *fiber.Ctx) error {
	db := config.GetDb()
	var stocks []models.Stock
	err := db.Find(&stocks).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch the stocks",
		})
	}
	return c.Status(200).JSON(stocks)
}

func GetStockById(c *fiber.Ctx) error {
	db := config.GetDb()
	var stock models.Stock
	stock_id := c.Params("id")
	err := db.Where("stock_id=?", stock_id).First(&stock).Error
	if err != nil {
		errmessage := fmt.Sprintf("Failed to fetch the stock with an id : %s", stock_id)
		return c.Status(404).JSON(fiber.Map{
			"error": errmessage,
		})
	}
	return c.Status(200).JSON(stock)
}

func CreateStock(c *fiber.Ctx) error {
	db := config.GetDb()
	var stock models.Stock
	err := utils.ParseData(&stock, c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err = db.Create(&stock).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create the stock",
		})
	}
	return c.Status(201).JSON(stock)
}

func UpdateStockById(c *fiber.Ctx) error {
	db := config.GetDb()
	stock_id := c.Params("id")
	var stock models.Stock
	err := db.Where("stock_id = ?", stock_id).First(&stock).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Stock not found",
		})
	}
	err = utils.ParseData(&stock, c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err = db.Save(&stock).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update stock",
		})
	}
	return c.JSON(stock)
}

func DeleteStockById(c *fiber.Ctx) error {
	db := config.GetDb()
	stock_id := c.Params("id")
	var stock models.Stock
	err := db.Where("stock_id = ?", stock_id).First(&stock).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Stock not found",
		})
	}
	err = db.Delete(&stock).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete stock",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Stock deleted successfully",
	})
}
