package cronjob

import (
	"github.com/hyc3z/Omaticaya/src/global"
	"github.com/hyc3z/Omaticaya/src/gpu"
	"github.com/hyc3z/Omaticaya/src/operation"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"sync"
)

func InitCronjob() *cron.Cron {
	return cron.New()
}

func UpdateGPUInfo(c *cron.Cron) {
	if err := c.AddFunc(global.ProjectInfo.MonitoringGpuIntervalPattern, func() {
		err := gpu.GetGpuInfo()
		if err != nil {
			global.Logger.Error("UpdateGPUInfo error", zap.Error(err))
			return
		}
		operation.UpdateTagForNode()
	}); err != nil {
		global.Logger.Error("UpdateGPUInfo error", zap.Error(err))
	}
}

func UpdateSchedulingPolicy(c *cron.Cron) {
	if err := c.AddFunc(global.ProjectInfo.MonitoringPolicyIntervalPattern, func() {

	}); err != nil {
		global.Logger.Error("UpdateSchedulingPolicy error", zap.Error(err))
	}
}

func StartJob(c *cron.Cron, w *sync.WaitGroup) {
	w.Add(1)
	c.Start()
}
