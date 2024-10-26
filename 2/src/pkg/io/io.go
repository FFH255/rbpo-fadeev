package io

import (
	"bufio"
	"strings"
)

func ReadLine(in *bufio.Reader) string {
	s, _ := in.ReadString('\n')
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
