package main

import (
	"errors"
	"regexp"
)

const urlRegexpStr = `^https?://.+`

func isUrlReal(url string) (bool, error) {
	if url == "" {
		return false, errors.New("url is empty")
	}

	var urlRegexp = regexp.MustCompile(urlRegexpStr)

	return urlRegexp.MatchString(url), nil
}
