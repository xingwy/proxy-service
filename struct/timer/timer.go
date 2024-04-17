package timer

import (
	"proxy-service/constants"
	"sync"
	"time"
)

type Timer struct {
	tasks      map[constants.TIMER_TYPE]*time.Ticker
	lock       sync.Mutex
	wg         sync.WaitGroup
	stopSignal chan bool
}

func NewTimer() *Timer {
	return &Timer{
		tasks:      make(map[constants.TIMER_TYPE]*time.Ticker),
		stopSignal: make(chan bool),
	}
}

func (t *Timer) AddTask(tt constants.TIMER_TYPE, duration time.Duration, task func()) {
	t.lock.Lock()
	defer t.lock.Unlock()

	// 创建一个定时器
	ticker := time.NewTicker(duration)

	// 启动一个协程来执行循环任务
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-ticker.C:
				task()
			case <-t.stopSignal:
				ticker.Stop()
				return
			}
		}
	}()

	t.tasks[tt] = ticker
}

func (t *Timer) RemoveTask(tt constants.TIMER_TYPE) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if ticker, ok := t.tasks[tt]; ok {
		ticker.Stop()
		delete(t.tasks, tt)
	}
}

func (t *Timer) Stop() {
	t.stopSignal <- true
	t.wg.Wait()
	close(t.stopSignal)
}
