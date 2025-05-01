package middleware

import (
	"cinemago/internal/model/dto"
	"github.com/gofiber/fiber/v2"
)

func NotFound(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "Not Found",
		})
	})
}
