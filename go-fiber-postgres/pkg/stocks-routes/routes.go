package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sai435603/go-fiber-postgres/pkg/stocks-controllers"
)

func RegisterStockRoutes(app *fiber.App) {
	app.Get("/api/stock", controllers.GetStocks)
	app.Get("/api/stock/:id", controllers.GetStockById)
	app.Post("/api/stock", controllers.CreateStock)
	app.Put("/api/stock/:id", controllers.UpdateStockById)
	app.Delete("/api/stock/:id", controllers.DeleteStockById)
}
