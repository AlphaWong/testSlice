package main

import (
	"strings"
)

func doJoinBySlice(input []string) string {
	return strings.Join(input[:], ",")
}

func doJoin(input []string) string {
	return strings.Join(input, ",")
}

func doJoinByCopy(input []string) string {
	tmpCopy := make([]string, len(input))
	copy(tmpCopy, input)
	return strings.Join(tmpCopy, ",")
}
