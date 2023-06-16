package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func LogError(err error) {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
	Logger.Error("Error: ", zap.Error(err))
}

func LogInfo(msg string) {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
	Logger.Info("Info: ", zap.String("message", msg))
}

func LogFatal(err error) {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync()
	Logger.Fatal("Fatal: ", zap.Error(err))
}