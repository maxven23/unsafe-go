package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x float64 = 100.0
	xPtr := unsafe.Pointer(&x)
	yPtr := (*int64)(xPtr)
	fmt.Println("y:", *yPtr)
}
