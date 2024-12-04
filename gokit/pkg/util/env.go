package util

import (
	"os"
	"strconv"
)

func BoolEnv(key string, def bool) bool {
	vv, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	v, err := strconv.ParseBool(vv)
	if err != nil {
		return def
	}

	return v
}
