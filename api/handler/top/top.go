package top

import (
	"proxy-service/struct/define"
	"proxy-service/struct/event"
)

type TopInstance struct {
	define.Instance
	EventManager event.EventManager
}

