package cron

// JobWrapper decorates the given Job with some behavior.
type TJobWrapper func(IJob) IJob
