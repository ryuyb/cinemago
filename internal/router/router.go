package router

import "github.com/gofiber/fiber/v2"

type Router interface {
	RegisterRoutes(app *fiber.App)
}

func SetupRoutes(app *fiber.App, routes []Router) {
	for _, route := range routes {
		route.RegisterRoutes(app)
	}
}
