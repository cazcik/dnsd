package main

import (
	"log"

	"github.com/cazcik/utils/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
