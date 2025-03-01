package main

import (
	"ocpi-cpo-mock-server/src/api/middlewares"
	"ocpi-cpo-mock-server/src/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	middlewares.RegisterMiddlewares(app)
	routes.RegisterRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
