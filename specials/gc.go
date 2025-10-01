package main

import (
	"runtime"
	"unsafe"
)

func main() {
	x := new(int)
	y := new(int)

	ptrX := unsafe.Pointer(x)
	addressY := uintptr(unsafe.Pointer(y))

	runtime.GC()

	*(*int)(ptrX) = 100
	*(*int)(unsafe.Pointer(addressY)) = 200
}
