package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var global bool
	var text string = "Lorem ipsum."
	var negative int = -9 // int8  int16  int32  int64
	var positive uint = 9 // uint8 uint16 uint32 uint64
	var address uintptr = uintptr(unsafe.Pointer(&global))
	var ascii byte = 'a'               // alias for uint8
	var unicode rune = 'Ā'             // alias for int32
	var pi float32 = 3.141592653589793 // float32 float64
	var fullPi float64 = float64(pi)
	var complicated complex64 = 1 + 1i // complex64 complex128
	inference := 1 + complicated

	fmt.Println(global)
	fmt.Println(text)
	fmt.Println(negative)
	fmt.Println(positive)
	fmt.Println(address)
	fmt.Println(ascii)
	fmt.Println(unicode)
	fmt.Println(pi)
	fmt.Println(fullPi)
	fmt.Println(complicated)
	fmt.Println(inference)
}
