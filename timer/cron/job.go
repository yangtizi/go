package cron

// Job 是提交的cron作业的接口。
type IJob interface {
	Run()
}
