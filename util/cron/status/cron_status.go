package status

import (
	"github.com/dreamlu/gt/tool/log"
	"github.com/robfig/cron/v3"
)

// CronStatus
// 定时任务库: https://github.com/robfig/cron
// 定时扫描删除待支付订单
func CronStatus() {
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc("@every 2m", updateStatus) // 每2分钟扫描
	c.Start()
}

// 定时检测状态修改
func updateStatus() {
	log.Info("[定时处理状态开始]")
	// 任务预警
	//gt.NewCrud().Select("update eg_issue set status = 2 where status = 0 and DATE_SUB(now(), INTERVAL -end_hour HOUR) > end_time").Exec()
	//gt.NewCrud().Select("update pr_work_order set status = 2 where status = 0 and DATE_SUB(now(), INTERVAL -end_hour HOUR) > end_time").Exec()
	//gt.NewCrud().Select("update eg_issue set status = 3 where status = 2 and now() > end_time").Exec()
	//gt.NewCrud().Select("update pr_work_order set status = 3 where status = 2 and now() > end_time").Exec()
}
