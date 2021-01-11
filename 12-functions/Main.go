package main

import "fmt"

// --- function synatx

// func [name]() {
//... code
//}
// func keyword
// naming convention the same as variables, scopes based on upper or lower case

// --- parameters

// func addAndDisplay(a int, b int) {
// 	fmt.Println(a, " + ", b, " = ", a+b)
// }

//syntactic sugar for same-type parameters
func addAndDisplay(a, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}

// --- we can pass arguments by reference (pointer)

// #####################################################
// --- return values
//
func sumAndReturn(a, b int) int {
	return a + b
}

// --- return multiple values
func divideAndReturn(a, b float64) (float64, error) {
	// error handling at the begining
	if b == 0.0 {
		return 0.0, fmt.Errorf("Pamiętaj..., nie dziel przez zero")
	}
	return a / b, nil
}

func main() {

	addAndDisplay(1, 2)

	fmt.Printf("%v + %v = %v\n", 3, 4, sumAndReturn(3, 4))

	// #########################################
	// --- anonymous functions
	// immediately-invoked function

	func() {
		fmt.Println("Print from anonymous function")
	}()

	// function as a variable
	f := func() {
		fmt.Println("function as a variable")
	}
	f()
	// longer syntax
	var divide2 func(float64, float64) (float64, error)
	divide2 = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Pamiętaj..., nie dziel przez zero")
		}
		return a / b, nil
	}

	d, err2 := divide2(5.0, 0.2)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("divide function as a variable : ", d)

	c := counter{
		a: 1,
		b: 2,
		c: 3,
		d: 4,
	}
	fmt.Println("Result of a counter method : ", c.count())

	// --- typical usage of error variable
	a, b := 5.0, 0.0
	res, err := divideAndReturn(a, b)
	// error handling
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a, "/", b, "=", res)
}

// #####################################################
// --- methods
// function that excutes in context of a type (e.g. struct)
type counter struct {
	a int
	b int
	c int
	d int
}

func (c counter) count() int { // c ounter is a reciever, copy is created for this function
	return c.a + c.b + c.c + c.d
}

// ! recevier can be passed by reference, so a method can change its original values

// ###############################
// --- function in go can return address of local variable
// when function returns the value does not vanish
// it is automatically promoted from stack to heap

//e.g.
func sumAndReturnPointer(a, b int) *int {
	result := a + b
	return &result // this variable will be accessible even after function execution
}
