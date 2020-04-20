package global

import (
	"github.com/hyc3z/Omaticaya/src/gpu"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

type NodeInfo struct {
	NodeName string
	HasGpu   bool
	Gpus     []gpu.GPU
	Config   *rest.Config
}

type GlobalInfo struct {
	ProjectName                     string
	Version                         string
	SchedulingPolicy                string
	MonitoringGpuIntervalPattern    string
	MonitoringPolicyIntervalPattern string
	Node                            NodeInfo
}

var ProjectInfo GlobalInfo

func (f *GlobalInfo) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ProjectName", f.ProjectName)
	enc.AddString("Version", f.Version)
	enc.AddString("NodeName", f.Node.NodeName)
	enc.AddString("Policy", f.SchedulingPolicy)
	enc.AddString("Monitoring Gpu Interval", f.MonitoringGpuIntervalPattern)
	enc.AddString("Monitoring New Policy Interval", f.MonitoringPolicyIntervalPattern)
	return nil
}

func InitInClusterConfig() error {
	// InClusterConfig
	Logger.Info("Init kubernetes Config. ")
	Config, err := rest.InClusterConfig()
	if err != nil {
		Logger.Error("Init Cluster Config Error",
			zap.Error(err))
		return err
	}
	ProjectInfo.Node.Config = Config
	return nil
}

func InitInfo() error {
	nodeName := os.Getenv("NODE_NAME")
	policy := os.Getenv("SCHEDULING_POLICY")
	gpuInterval := os.Getenv("MONITOR_GPU_INTERVAL_PATTERN")
	policyInterval := os.Getenv("MONITOR_POLICY_INTERVAL_PATTERN")
	ProjectInfo = GlobalInfo{
		ProjectName: "Omaticaya",
		Version:     "v1.0",
		Node: NodeInfo{
			NodeName: nodeName,
			HasGpu:   false,
		},
		SchedulingPolicy:                policy,
		MonitoringGpuIntervalPattern:    gpuInterval,
		MonitoringPolicyIntervalPattern: policyInterval,
	}
	Logger.Info("InitInfo",
		// Structured context as strongly typed Field values.
		zap.Object("globalInfo", &ProjectInfo),
	)
	err := InitInClusterConfig()
	if err != nil {
		return err
	}
	return nil
}
