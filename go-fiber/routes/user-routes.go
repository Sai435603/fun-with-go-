package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber/controllers"
)

func RegisterUserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
