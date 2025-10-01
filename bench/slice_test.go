package main

import (
	"testing"
	"unsafe"
)

var Result string

func ToStringDeprecated(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func ToStringNew(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(data), len(data))
}

func BenchmarkConversionBase(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = string(slice)
	}
}

func BenchmarkUnsafeConversionNew(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = ToStringNew(slice)
	}
}

func BenchmarkUnsafeConversionDeprecated(b *testing.B) {
	slice := []byte("Hello world!!!")
	for i := 0; i < b.N; i++ {
		Result = ToStringDeprecated(slice)
	}
}
