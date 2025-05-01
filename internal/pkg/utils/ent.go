package utils

import (
	"cinemago/internal/model/dto"
	"cinemago/internal/model/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/samber/lo"
)

// ParseEntError Parsing ent errors and converting them to ErrorResponse
func ParseEntError(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := lo.ErrorsAs[*ent.ConstraintError](err); ok {
		return &dto.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: e.Error(),
		}
	}
	if e, ok := lo.ErrorsAs[*ent.NotFoundError](err); ok {
		return &dto.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: e.Error(),
		}
	}
	if e, ok := lo.ErrorsAs[*ent.ValidationError](err); ok {
		return &dto.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: e.Error(),
		}
	}
	log.Errorf("Occured an not supportted ent error: %+v", err)
	return &dto.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
	}
}
