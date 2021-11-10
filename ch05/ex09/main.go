package main

import (
	"regexp"
)

var registerMatcher = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	if len(s) <= 1 {
		return s
	}
	if f == nil {
		return s
	}
	return registerMatcher.ReplaceAllStringFunc(s, func(register string) string {
		return f(register[1:])
	})
}
