package middleware

import (
	"cinemago/internal/model/dto"
	"github.com/gofiber/fiber/v2/log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// ErrorHandler is a middleware for capturing and handling errors in HTTP requests
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var resp *dto.ErrorResponse

	if e, ok := lo.ErrorsAs[*fiber.Error](err); ok {
		resp = &dto.ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
		}
	} else if e, ok := lo.ErrorsAs[*dto.ErrorResponse](err); ok {
		resp = &dto.ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
			Details: e.Details,
		}
	} else {
		resp = &dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorw("Request processing error",
			"method", ctx.Method(),
			"path", ctx.Path(),
			"status", http.StatusInternalServerError,
			"error", err.Error(),
		)
	}

	return ctx.Status(resp.Code).JSON(resp)
}
