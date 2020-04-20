package cronjob

import (
	"github.com/hyc3z/Omaticaya/src/global"
	"github.com/hyc3z/Omaticaya/src/gpu"
	"github.com/hyc3z/Omaticaya/src/operation"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"sync"
	"time"
)

func InitCronjob() *cron.Cron {
	return cron.New()
}

func UpdateGPUInfo(c *cron.Cron) {
	if err := c.AddFunc(global.ProjectInfo.MonitoringGpuIntervalPattern, func() {
		t0 := time.Now()
		err := gpu.GetGpuInfo()
		t1t0 := time.Since(t0)
		global.Logger.Info("GetGPUInfo Spent",
			zap.Int64("GetGPUInfo Duration", t1t0.Milliseconds()),
		)
		t1 := time.Now()
		if err != nil {
			global.Logger.Error("UpdateGPUInfo error", zap.Error(err))
			return
		}
		operation.UpdateTagForNode()
		t2t1 := time.Since(t1)
		global.Logger.Info("UpdateTag Spent",
			zap.Int64("UpdateTag Duration", t2t1.Milliseconds()),
		)
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
