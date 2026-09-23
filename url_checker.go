package main

import "regexp"

func isUrlReal(url string) bool {
	var urlRegexp = regexp.MustCompile(`^https?://.+`)

	return urlRegexp.MatchString(url)
}
