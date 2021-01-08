package main

import (
	"fmt"
)

func main() {
	fmt.Println("aa")

	// --- for loop syntax

	fmt.Println("Basic for loop")
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		// i is scoped to the loop
	}
	//we cannot access i outside the loop

	// --- multiple variables

	fmt.Println("loop with multiple variables")
	for i, j := 1, 2; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// --- increment operation is not an expression

	// --- while
	// for ; i<5 ; {}

	fmt.Println("while")
	i := 0
	for i < 5 { // while in golang ;D
		fmt.Println(i)
		i++
	}

	// infinite loop
	fmt.Println("infinite loop")
	for {
		fmt.Println(i)
		i--
		if i == 0 {
			break // stop the ifinite loop
			// continue // next step of a loop
		}
	}

	// --- labels
	// ? how to break outside nested loop

	fmt.Println("nested loop with label")

Loop_label: // set a label
	for a := 1; a < 4; a++ {
		for b := 1; b < 4; b++ {
			fmt.Println(a * b)
			if a*b > 3 {
				break Loop_label // break to a label
			}
		}
	}
	// --- loop over collection
	populacjaMiast := map[string]int{
		"Warszawa": 1790658,
		"Kraków":   779115,
		"Łódź":     679941,
	}

	fmt.Println("loop over collection")
	for k, v := range populacjaMiast {
		fmt.Println(k, v)
	}
	// we can even loop over string, and a channel

}
