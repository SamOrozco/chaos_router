package util

import "strings"

type StringMatcher func(haystack, needle string) bool

var Equal StringMatcher = func(haystack, needle string) bool {
	return haystack == needle
}

var Contains StringMatcher = func(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
