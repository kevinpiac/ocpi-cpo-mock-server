package routes

import (
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"ocpi-cpo-mock-server/src/core/modules/versions"

	"github.com/gofiber/fiber/v2"
)

func ListVersions(c *fiber.Ctx) error {
	uc := versions.NewListVersionsUsecase()
	result := uc.Execute(c.Locals("control").(*controls.Control))
	return c.JSON(result)
}
