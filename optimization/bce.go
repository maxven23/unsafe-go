package main

import "unsafe"

func fn1(data []int, check func(int) bool) []int {
	var idx = 0
	for _, value := range data {
		if check(value) {
			data[idx] = value // Found IsInBounds
			idx++
		}
	}

	return data[:idx] // Found IsSliceInBounds
}

func fn2(lhs, rhs []byte) {
	for idx := range min(len(lhs), len(rhs)) {
		_ = lhs[idx]
		_ = rhs[idx]
	}
}

func fn3(data [256]byte) {
	for idx := 0; idx < 128; idx++ {
		_ = data[2*idx]
	}
}

func cat1(str string) string {
	return string(str[0]) + // Found IsInBounds
		string(str[1]) + // Found IsInBounds
		string(str[2]) + // Found IsInBounds
		string(str[3]) // Found IsInBounds
}

func cat2(str string) string {
	return string(str[3]) + // Found IsInBounds
		string(str[2]) +
		string(str[1]) +
		string(str[0])
}

func sum1(values []int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += values[i] // Found IsInBounds
	}
	return sum
}

func sum2(values []int, size int) int {
	_ = values[size-1] // Found IsInBounds

	sum := 0
	for i := 0; i < size-1; i++ {
		sum += values[i]
	}
	return sum
}

func sum3(values []int, size int) int {
	sum := 0
	elemSize := unsafe.Sizeof(sum)
	start := unsafe.Pointer(unsafe.SliceData(values))
	for i := 0; i < size-1; i++ {
		sum += *(*int)(unsafe.Add(start, uintptr(i)*elemSize))
	}
	return sum
}

func main() {

}
