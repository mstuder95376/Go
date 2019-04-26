package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int 32 int 64
	// uint uint8 uint 16 uint 32 uint 64
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Brad"
	var age = 37
	const isCool = true
	// Shorthand method
	lname := "Ackley"

	fmt.Println(name, lname, "age", age, isCool)
	fmt.Printf("%T\n", isCool)

}
