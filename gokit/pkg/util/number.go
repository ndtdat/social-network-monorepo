package util

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func HexToUint64(hex string) (uint64, error) {
	value, err := HexToBigInt(hex)
	if err != nil {
		return 0, err
	}

	return value.Uint64(), nil
}

func HexToBigInt(hex string) (*big.Int, error) {
	value := strings.ReplaceAll(hex, "0x", "")
	bNum := new(big.Int)
	_, ok := bNum.SetString(value, 16)
	if !ok {
		return nil, fmt.Errorf("cannot parse hex for %x", hex)
	}

	return bNum, nil
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))

	return math.Round(val*ratio) / ratio
}

func Uint64Pointer(v uint64) *uint64 {
	return &v
}

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

func MustParseUint64Array(values []string) []uint64 {
	var results []uint64

	for _, v := range values {
		results = append(results, MustParseUint64(v))
	}

	return results
}

func Uint64ToString(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func MustParseInt(v string) int {
	return MustParseIntWithDefault(v, 0)
}

func ParseIntWithDefault(v string, d int) (int, error) {
	if v == "" {
		return d, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("cannot parse int for %s", v)
	}

	return i, nil
}

func MustParseIntWithDefault(v string, d int) int {
	i, err := ParseIntWithDefault(v, d)
	if err != nil {
		panic(fmt.Sprintf("cannot parse int for %s", v))
	}

	return i
}

func NullableUint64ToString(n *uint64) string {
	if n == nil {
		return ""
	}

	return Uint64ToString(*n)
}

func Int32ArrayToString(values []int32, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), delim), "[]")
}

func IntToString(n int) string {
	return strconv.FormatUint(uint64(n), 10)
}
