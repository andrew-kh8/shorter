package main

import (
	"testing"
)

func TestIsUrlRealRandomString(t *testing.T) {
	invalidUrl := "not real url"

	result := isUrlReal(invalidUrl)

	if result {
		t.Error("omg, it's an error")
	}
}

func TestIsUrlRealHttp(t *testing.T) {
	invalidUrl := "http://s"

	result := isUrlReal(invalidUrl)

	if !result {
		t.Error("omg, it's an error")
	}
}

func TestIsUrlRealHttps(t *testing.T) {
	invalidUrl := "https://s"

	result := isUrlReal(invalidUrl)

	if !result {
		t.Error("omg, it's an error")
	}
}

func TestIsUrlRealNoDomen(t *testing.T) {
	invalidUrl := "http://"

	result := isUrlReal(invalidUrl)

	if result {
		t.Error("omg, it's an error")
	}
}
