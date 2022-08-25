package cron

import (
	"sync"
	"time"
)

// Cron跟踪任意数量的条目，调用关联的func作为
// 由附表规定。它可以启动、停止，并且条目可以
// 运行时进行检查。
type TCron struct {
	entries  []*TEntry // 用于存放job指针对象的数组
	chain    TChain
	stop     chan struct{}      // 定制调度任务
	add      chan *TEntry       // 添加一个调度任务
	remove   chan TEntryID      // 移除 一个调度任务
	snapshot chan chan []TEntry // 正在运行中的调度任务
	running  bool               // 保证整个Cron对象只启动一次 和启动后其他chan正常
	// logger    Logger
	runningMu sync.Mutex      // 协程锁，确保执行安全
	location  *time.Location  // 时区
	parser    IScheduleParser // 解析参数
	nextID    TEntryID        // 下一个调度任务的id
	jobWaiter sync.WaitGroup  // 确保单一的调度任务执行完毕
}

// New returns a new Cron job runner, modified by the given options.
//
// Available Settings
//
//	Time Zone
//	  Description: The time zone in which schedules are interpreted
//	  Default:     time.Local
//
//	Parser
//	  Description: Parser converts cron spec strings into cron.Schedules.
//	  Default:     Accepts this spec: https://en.wikipedia.org/wiki/Cron
//
//	Chain
//	  Description: Wrap submitted jobs to customize behavior.
//	  Default:     A chain that recovers panics and logs them to stderr.
//
// See "cron.With*" to modify the default behavior.
func New(opts ...TOption) *TCron {
	c := &TCron{
		entries:   nil,
		chain:     NewChain(),
		add:       make(chan *TEntry),
		stop:      make(chan struct{}),
		snapshot:  make(chan chan []TEntry),
		remove:    make(chan TEntryID),
		running:   false,
		runningMu: sync.Mutex{},
		// logger:    DefaultLogger,
		location: time.Local,
		parser:   standardParser, //解析器
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
