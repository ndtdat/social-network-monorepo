package random

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/log"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/set"
	"math/big"
)

func Float32() float32 {
	max := int64(1<<53 - 1)

	return float32(Int64(0, max)) / float32(max)
}

func Float32Range(min, max float32) float32 {
	return min + Float32()*(max-min)
}

func Float64Range(min, max float64) float64 {
	return min + Float64()*(max-min)
}

func Float64() float64 {
	max := int64(1<<53 - 1)

	return float64(Int64(0, max)) / float64(max)
}

func Int64(min, max int64) int64 {
	bg := big.NewInt(max - min + 1)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("Error when generating int64 due to %v", err))

		return 0
	}

	return n.Int64() + min
}

func ArrayUint32(num, min, max int64, unique bool) []uint32 {
	var numbers []uint32
	numberSet := set.New[uint32]()

	for i := 0; i < int(num); i++ {
		valid := false
		for {
			if valid {
				break
			}

			rn := uint32(Int64(min, max))
			if numberSet.Contains(rn) && unique {
				valid = false

				continue
			}

			numberSet.Add(rn)
			numbers = append(numbers, rn)
			valid = true
		}
	}

	return numbers
}

func Uint16() uint16 {
	var bytes [2]uint8

	if _, err := rand.Read(bytes[:]); err != nil {
		log.Logger().Error(fmt.Sprintf("Error when generating uint16 due to %v", err))

		return 0
	}

	return binary.LittleEndian.Uint16(bytes[:])
}

func Int16() int16 {
	return int16(Uint16())
}

//nolint:predeclared,revive
func byte() uint8 {
	var bytes [1]uint8

	if _, err := rand.Read(bytes[:]); err != nil {
		log.Logger().Error(fmt.Sprintf("Error when generating byte due to %v", err))

		return 0
	}

	return bytes[0]
}

func Int8() int8 {
	return int8(byte())
}

func String(charset string, length int) string {
	nCharset := len(charset)
	s := make([]uint8, length)

	for i := range s {
		s[i] = charset[Int64(0, int64(nCharset-1))]
	}

	return string(s)
}
