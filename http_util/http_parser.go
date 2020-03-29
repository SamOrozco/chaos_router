package http_util

import (
	"bufio"
	"strings"
)

type HttpParser interface {
	ParsePath(tcpReq string) string
	ParseHeaders(tcpReq string) map[string]string
}

type TcpHttpParser struct {
}

func NewTcpHttpParser() HttpParser {
	return &TcpHttpParser{}
}

func (t TcpHttpParser) ParsePath(tcpReq string) string {
	lineScanner := bufio.NewScanner(strings.NewReader(tcpReq))
	
	// read first line
	if lineScanner.Scan() {
		firstLine := lineScanner.Text()
		segments := strings.Split(firstLine, " ")
		
		if len(segments) != 3 {
			return ""
		}
		return segments[1]
	}
	return ""
}

func (t TcpHttpParser) ParseHeaders(tcpReq string) map[string]string {
	lineScanner := bufio.NewScanner(strings.NewReader(tcpReq))
	resultMap := make(map[string]string, 0)
	
	// since we are parsing the headers we don't can about the first row
	// remove first row
	lineScanner.Scan()
	
	for lineScanner.Scan() {
		txt := lineScanner.Text()
		if txt == "" {
			return resultMap
		}
		segs := strings.SplitAfterN(txt, ":", 2)
		// remove last char or ":"
		headerKey := segs[0][0 : len(segs[0])-1]
		resultMap[headerKey] = strings.TrimSpace(segs[1])
	}
	return resultMap
}
