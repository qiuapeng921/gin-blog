package crontab

import (
	"gin-blog/app/crontab/cronfunc"
	"github.com/robfig/cron"
)

func InitCronTab() {
	cronTab := cron.New()
	_ = cronTab.AddFunc("0 */1 * * * *", cronfunc.HandleTest)
	cronTab.Start()
}
