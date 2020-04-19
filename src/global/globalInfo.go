package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type GlobalInfo struct {
	ProjectName                     string
	Version                         string
	NodeName                        string
	SchedulingPolicy                string
	MonitoringGpuIntervalPattern    string
	MonitoringPolicyIntervalPattern string
}

var ProjectInfo GlobalInfo

func (f *GlobalInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ProjectName", f.ProjectName)
	enc.AddString("Version", f.Version)
	enc.AddString("NodeName", f.NodeName)
	enc.AddString("Policy", f.SchedulingPolicy)
	enc.AddString("Monitoring Gpu Interval", f.MonitoringGpuIntervalPattern)
	enc.AddString("Monitoring New Policy Interval", f.MonitoringPolicyIntervalPattern)
	return nil
}

func InitInfo() error {
	nodeName := os.Getenv("NODE_NAME")
	policy := os.Getenv("SCHEDULING_POLICY")
	gpuInterval := os.Getenv("MONITOR_GPU_INTERVAL_PATTERN")
	policyInterval := os.Getenv("MONITOR_POLICY_INTERVAL_PATTERN")
	ProjectInfo = GlobalInfo{
		ProjectName:                     "Omaticaya",
		Version:                         "v1.0",
		NodeName:                        nodeName,
		SchedulingPolicy:                policy,
		MonitoringGpuIntervalPattern:    gpuInterval,
		MonitoringPolicyIntervalPattern: policyInterval,
	}
	Logger.Info("InitInfo",
		// Structured context as strongly typed Field values.
		zap.Object("globalInfo", &ProjectInfo),
	)

	return nil
}
