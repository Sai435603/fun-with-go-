package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber/config"
	"github.com/sai435603/go-fiber/models"
)

var db *[]models.User

func init() {
	db = config.GetDbInstance()
}

func GetUsers(c *fiber.Ctx) error {
	db := config.GetDbInstance()
	return c.JSON(*db)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	db := config.GetDbInstance()
	*db = append(*db, user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	db := config.GetDbInstance()

	for _, user := range *db {
		if user.ID == uint(id) {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	var updatedUser models.User
	err = c.BodyParser(&updatedUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db := config.GetDbInstance()
	for i, user := range *db {
		if user.ID == uint(id) {
			(*db)[i] = updatedUser
			return c.JSON(updatedUser)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	db := config.GetDbInstance()
	for i, user := range *db {
		if user.ID == uint(id) {
			(*db) = append((*db)[:i], (*db)[i+1:]...)
			return c.JSON(fiber.Map{
				"message": "User deleted",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "User not found",
	})
}
