package routes

import (
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"ocpi-cpo-mock-server/src/core/modules/credentials"
	"ocpi-cpo-mock-server/src/core/modules/versions"

	"github.com/gofiber/fiber/v2"
)

func RegisterCredentials(c *fiber.Ctx) error {
	uc := credentials.NewRegisterCredentialsUsecase()
	credentialTokenA := c.Locals("auth_token").(string)
	version := c.Params("version")

	result := uc.Execute(versions.VersionNumber(version), credentialTokenA, c.Locals("control").(*controls.Control))

	return c.JSON(result)
}
