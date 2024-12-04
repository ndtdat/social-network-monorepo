package util

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func ArrayNumberToAnyNumberType[A, B Number](values []A) []B {
	res := make([]B, 0, len(values))
	for _, value := range values {
		res = append(res, B(value))
	}

	return res
}

func NullAnyTypeToType[T any](value *T) T {
	if value == nil {
		var res T

		return res
	}

	return *value
}

func UnmarshalTypeToByte[T any](m T) []byte {
	mb, _ := json.Marshal(m)

	return mb
}

func ArrayToArrayString[T any](values []T, delim string) string {
	if len(values) == 0 {
		return "[]"
	}

	return strings.Join(strings.Fields(fmt.Sprint(values)), delim)
}

func ArrayToString[T any](values []T, delim string) string {
	if len(values) == 0 {
		return ""
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), delim), "[]")
}

//nolint:govet
func ArrayToStringV2[T any](values []T, delim string) string {
	if len(values) == 0 {
		return ""
	}

	var val []string
	for _, v := range values {
		switch any(values[0]).(type) {
		case string:
			val = append(val, fmt.Sprintf("'%s'", v))
		default:
			val = append(val, fmt.Sprintf("%v", v))
		}
	}

	return strings.Join(val, delim)
}
