package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/ocpi/cpo/versions", ListVersions)
	app.Get("/ocpi/:version/details", GetVersionDetails)

	app.Post("/ocpi/:version/credentials", RegisterCredentials)
}
