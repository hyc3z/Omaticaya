package main

import (
	"github.com/hyc3z/Omaticaya/src/cronjob"
	"github.com/hyc3z/Omaticaya/src/global"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	global.InitLogger()
	defer global.Logger.Sync()
	global.InitInfo()
	cron := cronjob.InitCronjob()
	defer cron.Stop()
	cronjob.UpdateGPUInfo(cron)
	cronjob.UpdateSchedulingPolicy(cron)
	wg.Wait()
}
