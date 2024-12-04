package util

import (
	"encoding/base64"
	"unicode"
)

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}

	return true
}

func PairsToMap(kps []string) map[string]string {
	result := map[string]string{}
	nPair := len(kps) / 2
	for i := 0; i < nPair; i++ {
		v := kps[i*2+1]
		if !isASCII(v) {
			bv, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				v = string(bv)
			}
		}

		result[kps[i*2]] = v
	}

	return result
}

func MapToPairs(meta map[string]string) []string {
	var metaVars []string
	for k, v := range meta {
		if !isASCII(v) {
			v = base64.StdEncoding.EncodeToString([]byte(v))
		}

		metaVars = append(metaVars, k, v)
	}

	return metaVars
}
