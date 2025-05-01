package logger

import "github.com/google/wire"

var LoggerSet = wire.NewSet(
	NewLogger,
)
