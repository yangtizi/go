package cron

type IScheduleParser interface {
	Parse(spec string) (ISchedule, error)
}
