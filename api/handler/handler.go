package handler

import (
	"proxy-service/api/handler/top"
	"proxy-service/conf"
	"proxy-service/struct/event"
)

var (
	Instance *Handler
)

type Handler struct {
	topInstance *top.TopInstance
}

func New(cfg *conf.Config) (s *Handler) {
	// 事件管理器
	eventInstance := event.NewEventManager()

	Instance = &Handler{
		topInstance: &top.TopInstance{
			EventManager: eventInstance,
		},
	}

	return Instance.Init()
}

func (s *Handler) Init() *Handler {
	s.topInstance.Init()
	return s
}

// 获取 TopHandler
func (s *Handler) TopHandler() top.TopHandler {
	return s.topInstance
}
