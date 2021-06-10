package cron

import (
	"demo/util/cron/status"
)

// 定时任务执行方式:
// 1.
// for{
//	time.Sleep(time.Millisecond* 100)
//	fmt.Println("Hello")
//}
// 2.
// for range time.Tick(time.Millisecond*100){
//	fmt.Println("Hello")
//}
// 3.
// c := time.Tick(5 * time.Second)
//for {
//	<- c
//	go f()
//}

// 定时任务
func Cron() {
	go status.CronStatus()
}
