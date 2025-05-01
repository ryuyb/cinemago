package handler

import (
	"cinemago/internal/model/dto"
	"cinemago/internal/pkg/utils"
	"cinemago/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

// Login user login
//
//	@Summary		user login
//	@Description	User login with username and password returns an access token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			loginReq	body		dto.UserLoginReq	true	"LoginInfo"
//	@Success		200			{object}	dto.UserLoginResp
//	@Failure		400			{object}	dto.ErrorResponse
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	loginReq := dto.UserLoginReq{}
	err := c.BodyParser(&loginReq)
	if err != nil {
		return dto.BadRequest("Invalid login information")
	}
	user, err := h.userService.CheckPassword(loginReq)
	if user == nil || err != nil {
		return dto.BadRequest("Incorrect username or password")
	}
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return dto.BadRequest("Generate access token failed")
	}
	return c.JSON(dto.UserLoginResp{AccessToken: accessToken})
}
