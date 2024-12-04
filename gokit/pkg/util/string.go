package util

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"regexp"
	"strings"

	"github.com/alioygur/godash"
)

var (
	matchFirstCap     = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap       = regexp.MustCompile("([a-z0-9])([A-Z])")
	CapitalAlphabet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
	Numbers           = "0123456789"
	Specials          = "!@#$%^&*()_+.,[]\\|"
)

func StringInSlice(target string, list []string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}

	return false
}

func SubStringInSlice(target string, list []string) bool {
	for _, str := range list {
		if strings.Contains(target, str) {
			return true
		}
	}

	return false
}

func SplitByWidth(str string, size int) []string {
	strLength := len(str)
	var splited []string
	var stop int
	for i := 0; i < strLength; i += size {
		stop = i + size
		if stop > strLength {
			stop = strLength
		}
		splited = append(splited, str[i:stop])
	}

	return splited
}

func SliceToLower(slice []string) []string {
	for i, s := range slice {
		slice[i] = strings.ToLower(s)
	}

	return slice
}

func ToUpperFirst(s string) string {
	if s == "" {
		return s
	}

	if len(s) == 1 {
		return strings.ToLower(string(s[0]))
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

// ToLowerSnakeCase the given string in snake-case format.
func ToLowerSnakeCase(s string) string {
	return strings.ToLower(godash.ToSnakeCase(s))
}

func RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b) //nolint:errcheck

	return hex.EncodeToString(b)
}

func ToLowerDashCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")

	return strings.ToLower(snake)
}

func ToLowerUnderscore(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func NullableStringToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func RandomStringWithSpecialChar(length int, chars string) string {
	b := make([]byte, length)
	nChars := int64(len(chars))
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(nChars))
		b[i] = chars[int(n.Int64())]
	}

	return string(b)
}
