package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int64
	var y string = "Hello world"
	var z = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("size x:", unsafe.Sizeof(x))
	fmt.Println("size y:", unsafe.Sizeof(y))
	fmt.Println("size z:", unsafe.Sizeof(z))
}
