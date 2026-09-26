package main

import (
	"testing"
)

func TestIsUrlRealRandomString(t *testing.T) {
	invalidUrl := "not real url"

	result, err := isUrlReal(invalidUrl)

	if result {
		t.Error("omg, it's an error")
	}
	if err != nil {
		t.Error("There's an ERROR")
	}
}

func TestIsUrlRealHttp(t *testing.T) {
	invalidUrl := "http://s"

	result, err := isUrlReal(invalidUrl)

	if !result {
		t.Error("omg, it's an error")
	}
	if err != nil {
		t.Error("There's an ERROR")
	}
}

func TestIsUrlRealHttps(t *testing.T) {
	invalidUrl := "https://s"

	result, err := isUrlReal(invalidUrl)

	if !result {
		t.Error("omg, it's an error")
	}
	if err != nil {
		t.Error("There's an ERROR")
	}
}

func TestIsUrlRealNoDomen(t *testing.T) {
	invalidUrl := "http://"

	result, err := isUrlReal(invalidUrl)

	if result {
		t.Error("omg, it's an error")
	}
	if err != nil {
		t.Error("There's an ERROR")
	}
}

func TestIsUrlRealEmpty(t *testing.T) {
	invalidUrl := ""

	result, err := isUrlReal(invalidUrl)

	if result {
		t.Error("omg, it's an error")
	}
	if err == nil {
		t.Error("There's an ERROR")
	}
}
