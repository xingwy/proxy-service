package bucket

import (
	"proxy-service/struct/linknode"
	"sync"
	"time"
)

type _BucketItem[T any] struct {
	Pool   []T
	SyncAt time.Time
}

func (b *_BucketItem[T]) PopTopN(n int) []T {
	b.SyncAt = time.Now()
	if len(b.Pool) >= n {
		topN := b.Pool[:n]
		b.Pool = b.Pool[n:]
		return topN
	}
	topN := b.Pool
	b.Pool = make([]T, 0)
	return topN
}

func (b *_BucketItem[T]) PopTopAll() []T {
	b.SyncAt = time.Now()
	topN := b.Pool
	b.Pool = make([]T, 0)
	return topN
}

type _TaskItem[T any] struct {
	Data []T
	Key  string
}

type Bucket[T any] struct {
	bucket_size int
	task_size   int
	task        *linknode.LinkedList[_TaskItem[T]]
	consumer    func(key string, data []T)
	mu          sync.Mutex
	run         sync.Mutex
	values      map[string]*_BucketItem[T]
	interval    int
	ticker      *time.Ticker
	quit        chan struct{}
}

func NewBucket[T any](bucket_size, duration int, consumer func(key string, data []T)) *Bucket[T] {
	_b := &Bucket[T]{
		bucket_size: bucket_size,
		task_size:   20,
		consumer:    consumer,
		task:        &linknode.LinkedList[_TaskItem[T]]{},
		values:      make(map[string]*_BucketItem[T], 0),
		interval:    duration,
		ticker:      time.NewTicker(time.Duration(duration) * time.Second),
		quit:        make(chan struct{}),
	}
	// 注册消费
	go _b.consume()
	// 注册定时事件
	go _b.register()

	return _b
}

// 阻塞方法-同步push
func (b *Bucket[T]) Add(key string, values []T) {
	if len(values) <= 0 {
		return
	}

	for !b.canAdd() {
		time.Sleep(time.Millisecond * 100)
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.values[key]
	if !ok {
		b.values[key] = &_BucketItem[T]{
			Pool:   make([]T, 0),
			SyncAt: time.Now(),
		}
	}
	b.values[key].Pool = append(b.values[key].Pool, values...)
	// 是否触发消费某个桶
	for len(b.values[key].Pool) > int(b.bucket_size) {
		data := b.values[key].PopTopN(int(b.bucket_size))
		t := &_TaskItem[T]{
			Data: data,
			Key:  key,
		}
		b.buildTask(t)
	}
}

func (b *Bucket[T]) buildTask(t *_TaskItem[T]) {
	b.task.Append(*t)
}

func (b *Bucket[T]) register() {
	for {
		select {
		case <-b.ticker.C:
			for key, v := range b.values {
				data := v.PopTopAll()
				if len(data) <= 0 {
					continue
				}

				t := &_TaskItem[T]{
					Data: data,
					Key:  key,
				}
				b.buildTask(t)
			}
		case <-b.quit:
			// 收到退出信号，退出循环
			return
		}
	}
}

// 通知消费
func (b *Bucket[T]) consume() {
	for {
		select {
		case <-b.quit:
			// 收到退出信号，退出循环
			return
		default:
			b.run.Lock()
			_task := b.task.Pop()
			if _task == nil {
				b.run.Unlock()
				time.Sleep(time.Second * 1)
				continue
			}

			b.run.Unlock()
			b.consumer(_task.Key, _task.Data)
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (b *Bucket[T]) canAdd() bool {
	return b.task.Length < b.task_size
}

func (b *Bucket[T]) BeforeClose() {
	// 组装剩余的数据
	for key, v := range b.values {
		data := v.PopTopAll()
		if len(data) <= 0 {
			continue
		}

		t := &_TaskItem[T]{
			Data: data,
			Key:  key,
		}
		b.buildTask(t)
	}

	// 检查
	_task := b.task.Pop()
	for _task != nil {
		b.consumer(_task.Key, _task.Data)
		time.Sleep(time.Millisecond * 10)
		_task = b.task.Pop()
	}
}

// 关闭
func (b *Bucket[T]) CloseAndWait() {
	// 停止计时器
	b.ticker.Stop()
	// 关闭管道
	close(b.quit)
	// 处理最后的数据
	b.BeforeClose()
}

func (b *Bucket[T]) CacheCount() int {
	count := 0
	for _, v := range b.values {
		count += len(v.Pool)
	}
	return count
}
