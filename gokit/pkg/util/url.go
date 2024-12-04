package util

import (
	"net/url"
)

func MustJoinPathAllowEmpty(base string, elem string) string {
	if elem == "" {
		return ""
	}

	u, _ := url.Parse(elem)
	if u != nil && u.Scheme != "" {
		return elem
	}

	res, err := url.JoinPath(base, elem)
	if err != nil {
		return ""
	}

	return res
}
