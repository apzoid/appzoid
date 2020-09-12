package routes

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func echo(ctx *fiber.Ctx) {
	ctx.JSON(&fiber.Map{"success": true, "method": "GET"})
}

// SetupRoute uses to Configure routes for the application
func SetupRoute(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/*", echo)
}
