package global

import "go.uber.org/zap/zapcore"

type globalInfo struct {
	projectName string
	version     string
}

func (f *globalInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("projectName", f.projectName)
	enc.AddString("version", f.version)
	return nil
}
