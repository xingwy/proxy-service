package handler

import (
	"proxy-service/conf"
)

var (
	Instance *Handler
)

type Handler struct {
}

func New(cfg *conf.Config) (s *Handler) {
	// 事件管理器
	// eventInstance := event.NewEventManager()

	return Instance.Init()
}

func (s *Handler) Init() *Handler {

	return s
}

// 获取 TaobaoHandler
