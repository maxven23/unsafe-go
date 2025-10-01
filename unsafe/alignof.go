package main

import (
	"fmt"
	"unsafe"
)

type AlignofData struct {
	x int64
	y int32
	z int16
}

func main() {
	var data AlignofData
	fmt.Println("align data:", unsafe.Alignof(data))
	fmt.Println("align data.x:", unsafe.Alignof(data.x))
	fmt.Println("align data.y:", unsafe.Alignof(data.y))
	fmt.Println("align data.z:", unsafe.Alignof(data.z))
}
