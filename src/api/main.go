package main

import (
	"ocpi-cpo-mock-server/src/api/middlewares"
	"ocpi-cpo-mock-server/src/api/routes"
	"ocpi-cpo-mock-server/src/core/modules/env"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env.ValidateEnv()
	app := fiber.New()
	middlewares.RegisterMiddlewares(app)
	routes.RegisterRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
