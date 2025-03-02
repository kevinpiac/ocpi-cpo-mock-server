package middlewares

import (
	"fmt"
	"ocpi-cpo-mock-server/src/core/modules/controls"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ResponseControlHeader string

const (
	ResponseDelayControlHeader          ResponseControlHeader = "X-Response-Latency"
	ResponseTypeControlHeader           ResponseControlHeader = "X-Response-Type"
	ResponseHttpStatusControlHeader     ResponseControlHeader = "X-Response-Http-Status"
	ResponseOCPIStatusCodeControlHeader ResponseControlHeader = "X-Response-OCPI-Status"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(AuthMiddleware)
	app.Use(ResponseControlMiddleware)
}

func parseResponseLatency(c *fiber.Ctx) int {
	delay := c.Get(string(ResponseDelayControlHeader))
	if len(delay) == 0 {
		return 0
	}
	delayInt, err := strconv.Atoi(delay)

	if err != nil {
		return 0
	}
	return delayInt
}

func parseResponseType(c *fiber.Ctx) controls.ResponseTypeControlValue {
	responseType := c.Get(string(ResponseTypeControlHeader))
	if len(responseType) == 0 {
		return controls.ResponseTypeControlValueNormal
	}

	switch controls.ResponseTypeControlValue(responseType) {
	case controls.ResponseTypeControlValueNormal, controls.ResponseTypeControlValueEmpty, controls.ResponseTypeControlValueError:
		fmt.Printf("Response type: %s\n", responseType)
		return controls.ResponseTypeControlValue(responseType)

	default:
		fmt.Printf("Unrecognized response type: %s. Defaulting to %s\n", responseType, controls.ResponseTypeControlValueNormal)
		return controls.ResponseTypeControlValueNormal
	}
}

func ResponseControlMiddleware(c *fiber.Ctx) error {

	responseType := parseResponseType(c)

	latency := parseResponseLatency(c)

	if latency > 0 {
		fmt.Printf("Adding %dms latency\n", latency)
		time.Sleep(time.Duration(latency) * time.Millisecond)
	}

	c.Locals("control", &controls.Control{
		ResponseType: responseType,
	})

	return c.Next()
}
