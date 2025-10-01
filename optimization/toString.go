package main

import (
	"fmt"
	"unsafe"
)

func ToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func main() {
	data := []byte("hello, world")
	str := ToString(data)
	fmt.Printf(str)
}
