package event

import (
	"context"
	"proxy-service/constants"
	"proxy-service/utils"
	"reflect"
	"runtime"
	"sync"
)

type Event struct {
	Type constants.EVENT_TYPE
	Data []any
}

type EventHandler func(ctx context.Context, data Event) error

// EventManager 是一个用于处理和分发事件的管理器。
type EventManagerInstance struct {
	mu       sync.RWMutex
	handlers map[constants.EVENT_TYPE]map[string]EventHandler
}

type EventManager interface {
	RegisterHandler(eventType constants.EVENT_TYPE, handler EventHandler)
	Emit(ctx context.Context, event Event)
	EmitSync(ctx context.Context, event Event) error
}

// NewEventManager 创建一个新的 EventManager。
func NewEventManager() *EventManagerInstance {
	return &EventManagerInstance{
		handlers: make(map[constants.EVENT_TYPE]map[string]EventHandler),
	}
}

// RegisterHandler 为特定事件类型注册事件处理程序。
// 设计handler，尽量保证函数幂等
// 参数：
//   - eventType：要为其注册处理程序的事件类型。
//   - handler：要注册的事件处理程序函数。
func (em *EventManagerInstance) RegisterHandler(eventType constants.EVENT_TYPE, handler EventHandler) {
	em.mu.Lock()
	defer em.mu.Unlock()

	if _, exist := em.handlers[eventType]; !exist {
		em.handlers[eventType] = make(map[string]EventHandler)
	}

	// 使用反射获取处理程序的函数名。
	fn := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

	// 使用函数名存储处理程序。
	em.handlers[eventType][fn] = handler
}

// Emit 异步发送事件给其处理程序。
// 参数：
//   - event：要分发给其已注册处理程序的事件。
func (em *EventManagerInstance) Emit(ctx context.Context, event Event) {

	handlers, found := em.handlers[event.Type]
	if found {
		for _, handler := range handlers {
			go func(_h EventHandler) {
				_e := _h(ctx, event)
				if _e != nil {
					utils.LogError(constants.LOG_ID__EVENT_HANDLE, event, _e)
				}
			}(handler)
		}
	}
}

// EmitSync 同步发送事件给其处理程序，并在任何处理程序失败时返回错误。
// 参数：
//   - event：要分发给其已注册处理程序的事件。
//
// 返回：
//   - error：指示事件处理失败的错误。
func (em *EventManagerInstance) EmitSync(ctx context.Context, event Event) error {

	handlers, found := em.handlers[event.Type]
	if found {
		for _, handler := range handlers {
			err := handler(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
