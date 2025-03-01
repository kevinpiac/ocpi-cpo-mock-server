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

func GetVersionDetails(c *fiber.Ctx) error {
	uc := versions.NewGetVersionDetailsUsecase()
	version := c.Params("version")

	versionNumber := versions.VersionNumber(version)

	result := uc.Execute(versionNumber, c.Locals("control").(*controls.Control))
	return c.JSON(result)
}
