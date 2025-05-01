package server

import (
	"cinemago/internal/config"
	"cinemago/internal/middleware"
	"cinemago/internal/router"
	pkgLogger "cinemago/pkg/logger"
	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
)

func NewFiberApp(routes []router.Router, logger *pkgLogger.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Cinemago",
		ErrorHandler: middleware.ErrorHandler,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	})

	log.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logger.Logger,
	}))

	middleware.FiberMiddleware(app, logger)

	router.SetupRoutes(app, routes)

	middleware.NotFound(app)

	return app
}

func StartServer(cfg *config.Config, app *fiber.App) {
	if err := app.Listen(cfg.Server.Addr); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}

func StartServerWithGracefulShutdown(cfg *config.Config, app *fiber.App) {
	idleConnectionsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Fatalf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnectionsClosed)
	}()

	StartServer(cfg, app)

	<-idleConnectionsClosed
}
