package cronjob

import (
	"github.com/hyc3z/Omaticaya/src/global"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"sync"
)

func InitCronjob() *cron.Cron {
	return cron.New()
}

func UpdateGPUInfo(c *cron.Cron) {
	name := global.ProjectInfo.NodeName
	if err := c.AddFunc(global.ProjectInfo.MonitoringGpuIntervalPattern, func() {
		global.Logger.Info(name)
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
