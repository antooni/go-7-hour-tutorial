package main

import (
	"fmt"
)

func main() {
	// --- primitive data types

	// boolean
	var b bool = false
	fmt.Printf("%v %T \n", b, b)

	// integers
	// int8 int16 int32 int64
	// uint8 uint16 uint32 uint64
	// default int is at least 32
	// var i int = 17

	// floating point numbers
	// float32 float64

	// strings
	// var s string = "this is a string"
	// ? string is an array of chars
	// e.g. s[2] (just int value on this position)
	// string(s[2]) (character on this position)
	// adding strings : s1 + s2

	// byte is uint8
	// string is a collection of bytes

	// --- bitwise operands
	// OR 	a | b
	// AND	a & b
	// XOR	a ^ b
	// LEFT SHIFT 	a << 3
	// RIGHT SHIFT 	a >> 3

	// --- you cannot implicit operation
	// var a int8 = 1
	// var b int16 = 5
	// e.g. fmt.Println(a + int8(b))
}
