package router

import (
	"github.com/cazcik/utils/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home",
		})
	})

	// Middleware
	api := app.Group("/api", logger.New())

	// Domain
	domain := api.Group("/domain")
	domain.Get("/:domain", handler.GetDomain)
}