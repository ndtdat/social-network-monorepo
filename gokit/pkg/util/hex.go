package util

import (
	"encoding/hex"
	"fmt"
)

func MustDecodeHexString(in string) []byte {
	bytes, err := hex.DecodeString(in)
	if err != nil {
		panic(fmt.Sprintf("cannot decode hex string due to %v", err))
	}

	return bytes
}
