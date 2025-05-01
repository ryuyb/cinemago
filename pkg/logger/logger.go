package logger

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.Logger to provide custom functionality
type Logger struct {
	*zap.Logger
}

// Config is the logger configuration struct
type Config struct {
	Level          string // Log level: debug, info, warn, error
	FilePath       string // Log file path
	MaxSize        int    // Maximum size of a single log file (MB)
	MaxBackups     int    // Maximum number of old log files to retain
	MaxAge         int    // Maximum number of days to retain old log files
	Compress       bool   // Whether to compress old log files
	Console        bool   // Whether to output to console
	File           bool   // Whether to output to file
	FileJsonFormat bool   // Whether to output to file in JSON format
	CallerSkip     int    // Increases the number of callers skipped by caller annotation
}

// defaultConfig returns the default logger configuration
func defaultConfig() Config {
	return Config{
		Level:          "info",
		FilePath:       "logs/app.log",
		MaxSize:        100,
		MaxBackups:     5,
		MaxAge:         30,
		Compress:       true,
		Console:        true,
		File:           false,
		FileJsonFormat: false,
		CallerSkip:     0,
	}
}

// NewLogger creates and initializes a new logger instance
func NewLogger(cfg Config) (*Logger, error) {
	// Parse log level
	var level zapcore.Level
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	var cores []zapcore.Core
	if cfg.Console {
		encoderConfig := getEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), level)
		cores = append(cores, consoleCore)
	}
	if cfg.File {
		encoderConfig := getEncoderConfig()
		var encoder zapcore.Encoder
		if cfg.FileJsonFormat {
			encoder = zapcore.NewJSONEncoder(encoderConfig)
		} else {
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		}
		fileCore := zapcore.NewCore(encoder, getLogWriter(cfg), level)
		cores = append(cores, fileCore)
	}

	// Create zap logger
	zapLogger := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(cfg.CallerSkip),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return &Logger{zapLogger}, nil
}

// NewDefaultLogger creates and initializes a new logger instance with default config
func NewDefaultLogger(opts ...Option) (*Logger, error) {
	cfg := defaultConfig()

	// Apply configuration options
	for _, opt := range opts {
		opt(&cfg)
	}

	return NewLogger(cfg)
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// getLogWriter configures the log file writer with rotation support
func getLogWriter(cfg Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	// Ensure log directory exists
	dir := filepath.Dir(cfg.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	return zapcore.AddSync(lumberJackLogger)
}

// Option is the configuration option function type
type Option func(*Config)

// WithLevel sets the log level
func WithLevel(level string) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithFilePath sets the log file path
func WithFilePath(path string) Option {
	return func(c *Config) {
		c.FilePath = path
	}
}

// WithMaxSize sets the maximum size of a single log file (MB)
func WithMaxSize(size int) Option {
	return func(c *Config) {
		c.MaxSize = size
	}
}

// WithMaxBackups sets the maximum number of old log files to retain
func WithMaxBackups(backups int) Option {
	return func(c *Config) {
		c.MaxBackups = backups
	}
}

// WithMaxAge sets the maximum number of days to retain old log files
func WithMaxAge(age int) Option {
	return func(c *Config) {
		c.MaxAge = age
	}
}

// WithCompress sets whether to compress old log files
func WithCompress(compress bool) Option {
	return func(c *Config) {
		c.Compress = compress
	}
}

// WithConsole sets whether to output to console
func WithConsole(console bool) Option {
	return func(c *Config) {
		c.Console = console
	}
}

// WithFile sets whether to output to file
func WithFile(file bool) Option {
	return func(c *Config) {
		c.File = file
	}
}

// WithFileJsonFormat sets whether to output to file in JSON format
func WithFileJsonFormat(json bool) Option {
	return func(c *Config) {
		c.FileJsonFormat = json
	}
}
