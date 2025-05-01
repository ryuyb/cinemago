package utils

import (
	"cinemago/internal/model/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

func ParseValidatorError(err error) error {
	if err == nil {
		return nil
	}
	e, ok := lo.ErrorsAs[validator.ValidationErrors](err)
	if !ok {
		log.Errorw("Current error is not validator error", "error", err)
		return err
	}
	details := make(map[string]string, len(e))
	for _, err := range e {
		details[err.Field()] = err.Tag()
	}
	return &dto.ErrorResponse{
		Code:    fiber.StatusBadRequest,
		Message: "validation failed",
		Details: details,
	}
}
