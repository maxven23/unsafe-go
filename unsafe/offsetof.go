package main

import (
	"fmt"
	"unsafe"
)

type OffsetofData struct {
	x int64
	y int32
	z int16
}

func main() {
	var data OffsetofData
	fmt.Println("offset x:", unsafe.Offsetof(data.x))
	fmt.Println("offset y:", unsafe.Offsetof(data.y))
	fmt.Println("offset z:", unsafe.Offsetof(data.z))
}
