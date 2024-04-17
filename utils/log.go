package utils

import (
	"fmt"
	"proxy-service/constants"
	"runtime"
	"strings"
)

func LogError(t constants.LOG_ID, data ...any) {
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, false)
	stackLines := Split(string(buf[:stackSize]), "\n")
	_, file, line, _ := runtime.Caller(1)

	lineInfo := ""
	for _, lineText := range stackLines {
		if strings.Contains(lineText, fmt.Sprintf("%s:%d", file, line)) {
			lineInfo += lineText
		}
	}
	content := map[string]any{
		"param1":   t,
		"callline": lineInfo,
	}
	if len(data) >= 1 {
		content["param2"] = fmt.Sprintf("%v", data[0])
	}
	if len(data) >= 2 {
		content["param3"] = fmt.Sprintf("%v", data[1])
	}
	if len(data) >= 3 {
		content["param4"] = fmt.Sprintf("%v", data[2])
	}

	fmt.Println(t, data)
}

func LogInfo(t constants.LOG_ID, data ...any) {
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, false)
	stackLines := Split(string(buf[:stackSize]), "\n")
	_, file, line, _ := runtime.Caller(1)

	lineInfo := ""
	for _, lineText := range stackLines {
		if strings.Contains(lineText, fmt.Sprintf("%s:%d", file, line)) {
			lineInfo += lineText
		}
	}
	content := map[string]any{
		"param1":   t,
		"callline": lineInfo,
	}
	if len(data) >= 1 {
		content["param2"] = fmt.Sprintf("%v", data[0])
	}
	if len(data) >= 2 {
		content["param3"] = fmt.Sprintf("%v", data[1])
	}
	if len(data) >= 3 {
		content["param4"] = fmt.Sprintf("%v", data[2])
	}

	fmt.Println(t, data)
}
