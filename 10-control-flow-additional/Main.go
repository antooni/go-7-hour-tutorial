package main

import (
	"fmt"
)

func main() {
	// defered function are called at the end of function before return
	// they are excuted in LIFO order

	defer fmt.Println("poczatek")
	defer fmt.Println("srodek")
	defer fmt.Println("koniec")
	// output = koniec\n srodek\n poczatek\n

	// --- defer usecase
	// associate opening and closing of a resource
	// in a short line-span
	// it is more readable and easier to analyse
	// e.g. openResource()
	// ....
	// defer closeResource()
	// be careful in loops !

	// --- defer takes arguments in the moment of declaration
	// e.g a:=1
	// defer fmt.Println(a)
	// a = 2
	// ! it will print 1 !

	// #####################################
	// --- panic
	// when exception happens, something really bad for an execution
	// we can stop the program and display info

	panic("AAAAAAAAAAAAAAAAAAAAAAAAAAA")

	// --- defered function will fire before panic

	//#############################
	// --- recover
	// only useful in defered functions
	// returns an error
	// function will stop, but higher in a stack can continue
}
