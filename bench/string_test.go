package main

import (
	"testing"
	"unsafe"
)

var Result string

func ToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func BenchmarkConversion(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = string(slice)
	}
}

func BenchmarkUnsafeConversion(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = ToString(slice)
	}
}
