package main

import (
	"github.com/ezedinff/appzoid/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger())
	routes.SetupRoute(app)
	app.Listen(5000)
}
