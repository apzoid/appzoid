package routes

import "github.com/gofiber/fiber"

func APIRoutes(app *fiber.App) {
	api := app.Group("/api")
	APIVersions(api)
}

func APIVersions(api fiber.Router) {
	v1Routes(api)
}
func v1Routes(api fiber.Router) {
	v1 := api.Group("/v1")
	v1.Get("/users", echo)
}
func echo(ctx *fiber.Ctx) {
	ctx.Send("Hello")
}
