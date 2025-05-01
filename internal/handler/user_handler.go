package handler

import (
	"cinemago/internal/model/dto"
	"cinemago/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var saveUserReq dto.CreateUserReq
	if err := ctx.BodyParser(&saveUserReq); err != nil {
		return dto.BadRequest("invalid user request")
	}
	err := u.userService.CreateUser(saveUserReq)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

func (u *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var saveUserReq dto.UpdateUserReq
	if err := ctx.BodyParser(&saveUserReq); err != nil {
		return dto.BadRequest("invalid user request")
	}
	err := u.userService.UpdateUser(saveUserReq)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}

// GetUserById handles GET /user/:id
//
//	@Summary		Get user by ID
//	@Description	Get user details by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int	true	"User ID"
//	@Success		200			{object}	dto.UserResp
//	@Failure		400,404,500	{object}	dto.ErrorResponse
//	@Router			/user/{id} [get]
func (u *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.BadRequest(fmt.Sprintf("Invalid id: %s", idStr))
	}
	user, err := u.userService.GetUserByID(id)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (u *UserHandler) DeleteUserById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.BadRequest(fmt.Sprintf("Invalid id: %s", idStr))
	}
	err = u.userService.DeleteUser(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}
