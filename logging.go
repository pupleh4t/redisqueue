package redisqueue

import (
	"os"
	"strings"
)

var (
	isDebug = false
)

func init() {
	d := os.Getenv("LIBQUEUE_DEBUG")
	if strings.ToLower(d) == "true" {
		isDebug = true
	}
}
