package main

import "unsafe"

func fn() {}

func main() {
	var str = "Hello world"
	const strCopy = "Hello world"

	println("address fn:     ", fn)
	println("address str:    ", unsafe.StringData(str))
	println("address strCopy:", unsafe.StringData(strCopy))
}
