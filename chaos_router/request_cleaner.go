package chaos_router

import (
	"fmt"
	"regexp"
	"strings"
)

type RequestCleaner interface {
	Clean(string) string
}

type NoOpCleaner struct {
}

func (NoOpCleaner) Clean(val string) string {
	return val
}

const hostHeaderPattern = "^Host.+[0-9]+$"

type HostHeaderCleaner struct {
	pattern     *regexp.Regexp
	routeToHost string
}

func NewHostHeaderCleaner(routeHost string) RequestCleaner {
	reg, err := regexp.Compile(hostHeaderPattern)
	if err != nil {
		return NoOpCleaner{}
	}
	return &HostHeaderCleaner{pattern: reg, routeToHost: routeHost}
}

func (h HostHeaderCleaner) Clean(req string) string {
	if found := h.pattern.FindString(req); found != "" {
		return strings.Replace(req, found, fmt.Sprintf("Host: %s", h.routeToHost), 1)
	}
	println("could not clean request")
	return req
}
