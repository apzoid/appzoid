package app

import (
	"github.com/ezedinff/appzoid/config"
	"github.com/ezedinff/appzoid/routes"
	"github.com/gofiber/fiber"
)

// App global variable for fiber app instance
var App *fiber.App

// Ctx global fiber context
var Ctx *fiber.Ctx

func bootApp() {
	App = fiber.New()
}
func Serve() {
	bootApp()
	config.LoadEnv()
	config.LoadAppConfig()
	routes.RegisterRoute(App)
	err := App.Listen(config.AppConfig.App_Port)
	if err != nil {
		panic("App not starting: " + err.Error() + "on Port: " + config.AppConfig.App_Port)
	}
}
