package constants

// 空值定义
var EMPTY_STR = ""
var EMPTY_TRUE = true
var EMPTY_FALSE = false
var MIN_DOUBLE float64 = 0.00000000001

// 事件类型
type EVENT_TYPE string

// 告警类型
type LOG_ID string

const (
	LOG_ID__EVENT_HANDLE LOG_ID = "EVENT_HANDLE"
	LOG_ID__COMMON       LOG_ID = "COMMON"
)

type TIMER_TYPE string
