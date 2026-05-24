package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber/routes"
)

func RootRoute(r *fiber.Ctx) error {
	return r.SendString("Hello World!")
}

func main() {
	app := fiber.New()
	app.Get("/", RootRoute)
	log.Println("Server is running on port 8000")
	routes.RegisterUserRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
