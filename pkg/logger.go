package logger

import (
	"context"
	slogzap "github.com/samber/slog-zap"
	"go.uber.org/zap"
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func init() {
	zapLogger, _ := zap.NewProduction()
	logger = slog.New(slogzap.Option{Level: slog.LevelInfo, Logger: zapLogger}.NewZapHandler())
}

func Logger() *slog.Logger {
	return logger
}
func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	logger.DebugContext(ctx, msg, args)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	logger.InfoContext(ctx, msg, args)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	logger.WarnContext(ctx, msg, args)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	logger.ErrorContext(ctx, msg, args)
}

func Fatal(msg string, args ...any) {
	logger.Error(msg, args)
	os.Exit(1)
}

func FatalContext(ctx context.Context, msg string, args ...any) {
	logger.ErrorContext(ctx, msg, args)
	os.Exit(1)
}
