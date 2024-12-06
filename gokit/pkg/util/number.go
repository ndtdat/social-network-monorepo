package util

import (
	"fmt"
	"strconv"
)

func SafeParseUint64WithDefault(v string, d uint64) uint64 {
	if v == "" {
		return d
	}

	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return d
	}

	return i
}

func MustParseUint64(v string) uint64 {
	return MustParseUint64WithDefault(v, 0)
}

func MustParseUint64WithDefault(v string, d uint64) uint64 {
	if v == "" {
		return d
	}

	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot parse uint64 for %s", v))
	}

	return i
}

func Uint64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func IntToString(n int) string {
	return strconv.FormatUint(uint64(n), 10)
}
