package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type GlobalInfo struct {
	ProjectName string
	Version     string
}

var ProjectInfo GlobalInfo

func (f *GlobalInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("projectName", f.ProjectName)
	enc.AddString("version", f.Version)
	return nil
}

func InitInfo() error {
	ProjectInfo = GlobalInfo{
		ProjectName: "Omaticaya",
		Version:     "v1.0",
	}
	Logger.Info("InitInfo",
		// Structured context as strongly typed Field values.
		zap.Object("globalInfo", &ProjectInfo),
	)

	return nil
}
