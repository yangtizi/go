package cron

import (
	"sync"
	"time"
)

// Cron跟踪任意数量的条目，调用关联的func作为
// 由附表规定。它可以启动、停止，并且条目可以
// 运行时进行检查。
type Cron struct {
	entries   []*TEntry
	chain     Chain
	stop      chan struct{}
	add       chan *TEntry
	remove    chan TEntryID
	snapshot  chan chan []TEntry
	running   bool
	logger    Logger
	runningMu sync.Mutex
	location  *time.Location
	parser    ScheduleParser
	nextID    TEntryID
	jobWaiter sync.WaitGroup
}
