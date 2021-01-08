package main

import (
	"fmt"
	"strconv"
)

// --- declaring at package scope
// var i int = 17
// cannot use := in package scope

// --- naming variable - scope
// package scope - lower case
// globally visible - upper case
// block scope - between curly braces {}

// --- naming conventions
// length should reflect lifespan, e.g. short name "i" for instantly used variable
// acronyms should be uppercase e.g. theHTTP

func main() {
	// --- declaring variables
	// var i int
	// i = 17

	// var i int = 17

	i := 17

	var (
		a int = 15
		b int = 18
	)
	// --- we cannot redeclare variable
	// --- all variables must be used

	// --- explicit conversions
	// golang would now risk loosing information
	// var f float32 = 17.7
	// var ii int = int(aa)

	// --- string conversion
	liczba := 42
	var napis string = string(liczba)
	var napis2 string = strconv.Itoa(liczba)
	fmt.Println("napis = ", napis, "napis2 = ", napis2) // napis = * napis2 = 42

	fmt.Printf("%#v %T %v %T %v %T \n", i, i, a, a, b, b)

}
