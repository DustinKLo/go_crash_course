package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint 8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var name = "dustin lo"
	// name := "brad"
	name, email := "Brad", "dustinlo@gmail.com"
	size := 1.3
	var age int32 = 37
	const isCool = true

	fmt.Println(name, age, size, isCool, email)
	fmt.Printf("%T %T %T %T %T \n", name, age, size, isCool, email)
	fmt.Println(isCool)
}
