package util

import (
	"time"
)

// Now

func CurrentUnix() int64 {
	return time.Now().Unix()
}
