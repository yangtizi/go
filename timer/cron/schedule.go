package cron

import "time"

// Schedule 描述作业的工作周期。
type ISchedule interface {
	//Next返回下一个激活时间，晚于给定时间。
	//Next最初调用，然后每次运行作业时调用。
	Next(time.Time) time.Time
}
