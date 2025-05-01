package router

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type SwaggerRouter struct {
}

func NewSwaggerRouter() *SwaggerRouter {
	return &SwaggerRouter{}
}

func (r *SwaggerRouter) RegisterRoutes(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
