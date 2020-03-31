package util

import (
	"regexp"
	"strings"
)

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

var Prefix StringMatcher = func(haystack, needle string) bool {
	return strings.HasPrefix(haystack, needle)
}

var Suffix StringMatcher = func(haystack, needle string) bool {
	return strings.HasSuffix(haystack, needle)
}

var Regex StringMatcher = func(haystack, needle string) bool {
	if found, err := regexp.MatchString(needle, haystack); err != nil {
		return false
	} else {
		return found
	}
}
