package util

import "strings"

type StringMatcher func(haystack, needle string) bool

var All StringMatcher = func(haystack, needle string) bool {
	return true
}

var Equal StringMatcher = func(haystack, needle string) bool {
	return haystack == needle
}

var Contains StringMatcher = func(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
