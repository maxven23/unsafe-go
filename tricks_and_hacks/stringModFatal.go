package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str = "Hello world"
	strData := unsafe.StringData(str)

	slice := unsafe.Slice(strData, len(str))
	slice[0] = 'G'

	fmt.Println(str)
}
