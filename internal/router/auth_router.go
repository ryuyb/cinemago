package router

import (
	"cinemago/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	authHandler *handler.AuthHandler
}

func NewAuthRouter(authHandler *handler.AuthHandler) *AuthRouter {
	return &AuthRouter{authHandler: authHandler}
}

func (r *AuthRouter) RegisterRoutes(app *fiber.App) {
	group := app.Group("/api/auth")

	group.Post("/login", r.authHandler.Login)
}
