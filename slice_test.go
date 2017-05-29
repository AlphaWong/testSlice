package main

import (
	"strconv"
	"testing"
)

func BenchmarkDoJoinBySlice(b *testing.B) {
	str := getRamdomStringArray()
	for i := 0; i < b.N; i++ {
		doJoinBySlice(str)
	}
}

func BenchmarkDoJoin(b *testing.B) {
	str := getRamdomStringArray()
	for i := 0; i < b.N; i++ {
		doJoin(str)
	}
}

func BenchmarkDoJoinByCopy(b *testing.B) {
	str := getRamdomStringArray()
	for i := 0; i < b.N; i++ {
		doJoinByCopy(str)
	}
}

func getRamdomStringArray() []string {
	array := make([]string, 99999999)
	for i := 0; i < 99999999; i++ {
		array[i] = strconv.Itoa(i)
	}
	return array
}
