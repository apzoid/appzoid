package routes

import "github.com/gofiber/fiber"

func RegisterRoute(app *fiber.App) {
	APIRoutes(app)
	AuthRoutes(app)
}
