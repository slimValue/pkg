package logger

import (
	"go.uber.org/zap"

	"skydoor-policy-server/internal/config"
	"skydoor-policy-server/internal/types"
)

var (
	defaultLogger Logger
	zapLogger     *zap.Logger
)

func Init() {
	if config.GetEnv() == types.EnvRelease {
		zapLogger, _ = zap.NewProduction()
	} else {
		zapLogger, _ = zap.NewDevelopment()
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(1))
	defaultLogger = zapLogger.Sugar()
}

func Cleanup() {
	_ = zapLogger.Sync()
}

func GetLogger() *zap.Logger {
	return zapLogger
}
