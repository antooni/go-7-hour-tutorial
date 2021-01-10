package main

import "fmt"

func main() {
	// --- declaration of a pointer
	var a int = 17
	var ptrA *int = &a

	fmt.Println(ptrA)

	// --- dereferencing a pointer
	fmt.Println(ptrA, " = ", *ptrA)

	*ptrA = 18
	fmt.Println(a, " = ", *ptrA)

	// --- ! there is no pointer arithmetics !
	// for simplicity

	// --- pointers and structs
	var ms *myStruct

	// what is the value of uninitialised pointer ?
	fmt.Println(ms) // <nil>

	// --- when we use new keyword, struct is initialised as an empty object
	ms = new(myStruct)
	// other way around : ms = &myStruct{a:1}
	fmt.Println(ms)

	// --- how to access field through a pointer
	(*ms).a = 15
	fmt.Println(ms)
	// syntactic sugar of golang
	// compiler helps us out, so we can reference field just with a dot
	ms.a = 16
	fmt.Println(ms)

	// --- datatypes with underlined pointers, copy is made by refernce
	// slice, map

}

type myStruct struct {
	a int
}
