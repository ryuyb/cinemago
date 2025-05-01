package router

import (
	"cinemago/internal/handler"
	"cinemago/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{userHandler: userHandler}
}

func (r *UserRouter) RegisterRoutes(app *fiber.App) {
	group := app.Group("/api/user")

	group.Post("/", r.userHandler.CreateUser)
	group.Put("/", r.userHandler.UpdateUser)
	group.Get("/:id", middleware.JwtProtected(), r.userHandler.GetUserById)
	group.Delete("/:id", r.userHandler.DeleteUserById)
}
