package middleware

import (
	"cinemago/internal/config"
	"cinemago/internal/model/dto"
	"cinemago/internal/pkg/utils"
	"cinemago/internal/repository"
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

const (
	JwtContextKey  = "jwt"
	UserContextKey = "user"
)

func JwtProtected() fiber.Handler {
	jwtCfg := config.GetConfig().Jwt
	wareConfig := jwtMiddleware.Config{
		SigningKey:     jwtMiddleware.SigningKey{Key: []byte(jwtCfg.SigningKey)},
		ContextKey:     JwtContextKey,
		Claims:         &utils.JwtClaims{},
		SuccessHandler: jwtSuccessHandler,
		ErrorHandler:   jwtErrorHandler,
	}
	return jwtMiddleware.New(wareConfig)
}

func jwtSuccessHandler(c *fiber.Ctx) error {
	token, ok := c.Context().UserValue(JwtContextKey).(*jwt.Token)
	if !ok {
		return dto.NewErrorResponse(fiber.StatusInternalServerError, "Failed to parse token")
	}
	if !token.Valid {
		return dto.NewErrorResponse(fiber.StatusUnauthorized, "Invalid token")
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		log.Errorf("Failed to parse token subject: %v", err)
		return dto.NewErrorResponse(fiber.StatusInternalServerError, "Failed to parse token")
	}
	userId, err := strconv.Atoi(subject)
	if err != nil {
		log.Errorf("Failed to parse token subject: %v", err)
		return dto.NewErrorResponse(fiber.StatusInternalServerError, "Failed to parse token subject")
	}
	user, err := repository.GetClient().User.Get(c.Context(), userId)
	if err != nil {
		return dto.NewErrorResponse(fiber.StatusUnauthorized, "Invalid token")
	}
	c.Context().SetUserValue(UserContextKey, user)
	return c.Next()
}

func jwtErrorHandler(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return dto.BadRequest(err.Error())
	}
	return dto.NewErrorResponse(fiber.StatusUnauthorized, err.Error())
}
