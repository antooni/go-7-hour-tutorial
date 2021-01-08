package main

import (
	"fmt"
)

func main() {
	// #############################################
	// --- declaring an array

	arr := [3]int{17, 55, 1}
	// or 	arr := [...]int{17, 55, 1} // it will create arr just the size of arguemnts

	var arr2 [3]string

	fmt.Printf("Array: %v \n\n", arr)

	arr2[0] = "napis" // access and index
	n := len(arr2)    // get length of an array

	fmt.Printf("Array2: %v \nLength: %v \n\n", arr2, n)

	// ! array is a value in Go !
	// It is copied by value
	// how to solve it?  pointer

	a := [...]int{3, 2, 1}
	b := &a

	b[1] = 22
	fmt.Printf("a = %v \nb = %v\n\n", a, b)

	// ##############################################
	// --- declaring a slice

	slice := []int{1, 2, 3}

	slice2 := make([]int, 3, 100)
	// slices are referenced datatypes

	// --- appending slices

	slice2 = append(slice2, 1)
	slice2 = append(slice2, slice...) // spread operator is needed
	// append only accepts args the same type as slice

	// --- shifting

	slice2 = slice2[1:]             // remove first element
	slice2 = slice2[:len(slice2)-1] // remove last element
	// middle slice2 := append(a[:2],a[3:]...) // remove 3rd element

	//fmt.Printf("slice: %v %T \n", slice, slice)

	fmt.Printf("slice2: %v %T \n", slice2, slice2)
	fmt.Printf("length = %v\n", len(slice2))
	fmt.Printf("capacity = %v\n", cap(slice2))

}
