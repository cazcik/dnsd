package main

import (
	"log"

	"github.com/cazcik/dnsd/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("dist", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		AllowCredentials: true,
	}))

	api := app.Group("/api", logger.New())
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": nil, "status": "success", "data": nil, })
	})
	api.Post("/lookup", handler.Lookup)

	app.Static("/", "dist")
	app.Static("*", "./dist/index.html")

	log.Fatal(app.Listen(":8080"))
}