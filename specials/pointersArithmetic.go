package main

import (
	"fmt"
	"unsafe"
)

type PointerData struct {
	a int32
	b int32
	c int32
}

func main() {
	d := PointerData{a: 10, b: 20, c: 30}
	ptr := unsafe.Pointer(&d)

	offset := unsafe.Offsetof(d.b)
	ptrB := unsafe.Pointer(uintptr(ptr) + offset)

	valueB := *(*int32)(ptrB)

	offsetC := unsafe.Offsetof(d.c)
	ptrC := unsafe.Pointer(uintptr(ptr) + offsetC)
	valueC := *(*int32)(ptrC)

	fmt.Println("a:", d.a)
	fmt.Println("b через pointer arithmetic:", valueB)
	fmt.Println("c через pointer arithmetic:", valueC)

	x := [5]int32{100, 200, 300, 400, 500}
	i := 3

	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + uintptr(i)*unsafe.Sizeof(x[0]))
	valueXi := *(*int32)(e)

	fmt.Println("x[3] через unsafe.Pointer:", valueXi)
}
