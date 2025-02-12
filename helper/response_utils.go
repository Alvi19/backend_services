package helper

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type SuccessResponseData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(&SuccessResponseData{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, message string, data interface{}) error {
	var isDebug = os.Getenv("DEBUG")
	data2 := data

	if isDebug == "false" {
		data2 = nil
	}

	return c.Status(fiber.StatusInternalServerError).JSON(&SuccessResponseData{
		Status:  false,
		Message: message,
		Data:    data2,
	})
}
