package routes

import (
	"github.com/ezedinff/appzoid/controllers"
	"github.com/gofiber/fiber"
)

// AuthRoutes auth routes
func AuthRoutes(app *fiber.App) {
	app.Get("/login", controllers.Login)
}
