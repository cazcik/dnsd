package router

import (
	"github.com/cazcik/utils/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Domain
	domain := api.Group("/domain")
	domain.Get("/:domain", handler.GetDomain)
}