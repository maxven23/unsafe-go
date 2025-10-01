package main

import (
	"fmt"
	"unsafe"
)

type Union struct {
	value [8]byte
}

func (u *Union) SetInt64(value int64) {
	*(*int64)(unsafe.Pointer(&u.value)) = value
}

func (u *Union) GetInt64() int64 {
	return *(*int64)(unsafe.Pointer(&u.value))
}

func (u *Union) SetFloat64(value float64) {
	*(*float64)(unsafe.Pointer(&u.value)) = value
}

func (u *Union) GetFloat64() float64 {
	return *(*float64)(unsafe.Pointer(&u.value))
}

func main() {
	var u Union

	u.SetInt64(200)
	fmt.Printf("After SetInt64:\n")
	fmt.Printf("  GetInt64()  = 0x%016x\n", u.GetInt64())
	fmt.Printf("  GetFloat64()= %v  (интерпретация тех же байт как float64)\n\n", u.GetFloat64())

	u.SetFloat64(100.0)
	fmt.Printf("After SetFloat64(100.0):\n")
	fmt.Printf("  GetFloat64()= %v\n", u.GetFloat64())
	fmt.Printf("  GetInt64()  = 0x%016x  (битовое представление float64)\n\n", u.GetInt64())
}
