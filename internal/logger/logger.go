package logger

import (
	"cinemago/internal/config"
	"fmt"

	pkgLogger "cinemago/pkg/logger"
)

var (
	Logger *pkgLogger.Logger
)

func NewLogger(cfg *config.Config) (*pkgLogger.Logger, func(), error) {
	logCfg := cfg.Log

	Logger, err := pkgLogger.NewLogger(pkgLogger.Config{
		Level:          logCfg.Level,
		FilePath:       logCfg.File.FilePath,
		MaxSize:        logCfg.File.MaxSize,
		MaxBackups:     logCfg.File.MaxBackups,
		MaxAge:         logCfg.File.MaxAge,
		Compress:       logCfg.File.Compress,
		Console:        logCfg.EnableConsole,
		File:           logCfg.File.Enable,
		FileJsonFormat: logCfg.File.JsonFormat,
		CallerSkip:     logCfg.CallerSkip,
	})

	cleanup := func() {
		err := Logger.Sync()
		if err != nil {
			fmt.Printf("Error syncing logger: %v\n", err)
		}
	}

	return Logger, cleanup, err
}
