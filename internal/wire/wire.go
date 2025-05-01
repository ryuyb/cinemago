//go:build wireinject
// +build wireinject

package wire

import (
	"cinemago/internal/config"
	"cinemago/internal/handler"
	"cinemago/internal/logger"
	"cinemago/internal/repository"
	"cinemago/internal/router"
	"cinemago/internal/server"
	"cinemago/internal/service"
	pkgLogger "cinemago/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

type App struct {
	App    *fiber.App
	Config *config.Config
	Logger *pkgLogger.Logger
}

func NewApp(app *fiber.App, cfg *config.Config) *App {
	return &App{App: app, Config: cfg}
}

func InitializeApp(cmd *cobra.Command) (*App, func(), error) {
	panic(wire.Build(
		config.ConfigSet,
		logger.LoggerSet,
		repository.RepositorySet,
		service.ServiceSet,
		handler.HandlerSet,
		router.RouterSet,
		server.ServerSet,
		NewApp,
	))
}
