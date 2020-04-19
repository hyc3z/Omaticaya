package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/hyc3z/omaticaya/src/global"
)

var logger *zap.Logger



func InitLogger() {
	logger, _ = zap.NewProduction()
}

func main() {
	InitLogger()
	defer logger.Sync()
	info := globalInfo{
		projectName: "Omaticaya",
		version:     "v1.0",
	}
	logger.Info("Init",
		// Structured context as strongly typed Field values.
		zap.Object("globalInfo",&info),
	)
}


