package sysutils

import "time"

type ITimerCallback interface {
	OnTimer() error
}

// 定时器
type TTimer struct {
	duInterval time.Duration // 时间间隔
	chStop     chan bool     // 停止信号
	bEnabled   bool          // 是否启动
}

// 新定时器
func NewTimer(d time.Duration, sender ITimerCallback) *TTimer {
	p := &TTimer{
		duInterval: d,
		chStop:     make(chan bool),
		bEnabled:   false,
	}

	// 启动定时器
	go p.Run(sender)

	return p
}

// 手动运行
func (m *TTimer) Run(sender ITimerCallback) {
	if m.bEnabled {
		return
	}
	m.bEnabled = true

	ticker := time.NewTicker(m.duInterval)
	for {
		select {
		case <-ticker.C: // 到时间, 定期清理
			err := sender.OnTimer()
			if err != nil {
				ticker.Stop()
				return
			}

		case <-m.chStop: // 删除计时器
			ticker.Stop()
			return
		}
	}
}

// 停止定时器
func (m *TTimer) Stop() {
	m.chStop <- true
	m.bEnabled = false
}
