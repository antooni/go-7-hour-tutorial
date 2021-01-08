package main

import (
	"fmt"
)

func main() {
	// --- naming constants - scoping
	// uppercase - global e.g.const GlobalConst
	// lowercse - package e.g. const packageConst

	const myConst int = 17

	//  --- of course we cannot change value of constant
	// myConst = 11 - ERR

	//  --- const value has to be available at compile time
	// e.g. this would not work    const sinus float64 = math.Sin(1.14)

	//  --- constants shadowing is possible
	// not only by a value but even a type
	// e.g. globally : 		const a int = 17
	//		inside block : 	const a float64 = 1.7

	// --- enumarted constant

	const (
		// --- iota is scoped to a constant block
		// is is like IDs for variables

		//  --- very neat trick
		// it avoids the situation when default declaration ov variable points to a use case
		// which is not what we want for our code to be readable
		errZawod = iota // 0

		//  --- another possible trick for ^above^ situation
		// _ = iota
		// _ (underscore) is a special "write-only" variable
		// work basically like a garbage for values

		// ---  we can add an offset
		// errZawod = iota + 5 // 5
		// and next values will be 6 7 8 ...

		//  --- if we set iota next variables will be automatically enumerated
		hydraulik // 1
		elektryk  // 2
		mechanik  // 3

		// later we can switch based on the value
		// or whatever
	)

	fmt.Printf("%v %T\n", myConst, myConst)
}
