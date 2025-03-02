package middlewares

import (
	"encoding/base64"
	"log"
	"ocpi-cpo-mock-server/src/core/modules/response"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationErrorResponse = response.BaseResponse[string]

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Next()
	}

	token := strings.TrimPrefix(authHeader, "Token ")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthenticationErrorResponse{
			Data:          "",
			StatusCode:    response.StatusCodeGenericClientError,
			StatusMessage: "The authorization header is malformed it should be in the format 'Token <token>' as specified in the OCPI protocol.",
			TimeStamp:     time.Now().Format(time.RFC3339),
		})
	}

	decoded, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthenticationErrorResponse{
			Data:          "",
			StatusCode:    response.StatusCodeGenericClientError,
			StatusMessage: "The authorization header is malformed it should be in the format 'Token <token>' as specified in the OCPI protocol. Also the token should be base64 encoded.",
			TimeStamp:     time.Now().Format(time.RFC3339),
		})
	}

	decodedString := string(decoded)

	c.Locals("auth_token", decodedString)

	log.Println("Auth token: ", decodedString)

	return c.Next()
}
