package main

import (
	"github.com/hyc3z/Omaticaya/src/cronjob"
	"github.com/hyc3z/Omaticaya/src/global"
	"github.com/hyc3z/Omaticaya/src/operation"
	"go.uber.org/zap"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	global.InitLogger()
	defer func() {
		err := global.Logger.Sync()
		if err != nil {
			panic("Logger sync failed.")
		}
	}()
	err := global.InitInfo()
	if err != nil {
		global.Logger.Panic("InitInfo failed.",
			zap.Error(err),
		)
	}
	operation.CleanTag()
	cron := cronjob.InitCronjob()
	cronjob.UpdateGPUInfo(cron)
	cronjob.UpdateSchedulingPolicy(cron)
	defer operation.CleanTag()
	defer cron.Stop()
	cronjob.StartJob(cron, &wg)
	wg.Wait()
}
