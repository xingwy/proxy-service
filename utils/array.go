package utils

import (
	"strings"
)

func Split(src, f string) []string {
	if src == "" {
		return []string{}
	}
	return strings.Split(src, f)
}
