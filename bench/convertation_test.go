package main

import (
	"testing"
	"unsafe"
)

type Int int

var convertedData = make([]int, 1024)

func BenchmarkCast(b *testing.B) {
	data := make([]Int, 1024)
	for i := 0; i < b.N; i++ {
		for idx, value := range data {
			convertedData[idx] = int(value)
		}
	}
}

func BenchmarkUnsafeCast(b *testing.B) {
	data := make([]Int, 1024)
	for i := 0; i < b.N; i++ {
		convertedData = *(*[]int)(unsafe.Pointer(&data))
	}
}
