package main

import (
	"github.com/yashptel/go-api-template/pkg/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := initLogger()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	http.RunHttpServer()
}

func initLogger() *zap.Logger {

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
