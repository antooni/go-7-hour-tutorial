package main

import (
	"fmt"
)

func main() {
	populacjaMiast := map[string]int{
		"Warszawa": 1790658,
		"Kraków":   779115,
		"Łódź":     679941,
	}

	// no single line if statements (finally)

	// --- classic syntax
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// --- initialiser syntax
	// allows us to evaluate certain values inline
	// if [initialiser]; condition {}
	if pop, ok := populacjaMiast["Warszawa"]; ok {
		fmt.Println("Warszawa", pop)
	}

	// --- there is no parenthesis
	// if a < b {}

	// --- short-circuting
	// OR statement
	// if it evaluates first part as true, there is no need to evaluate to second
	// e.g. true || codeToDestroyHumanity()
	// second part is skipped (for greater good)

	// AND works similiary, but the opposite happens
	// false && codeToDestroyHumanity()
	// second part is skipped

	// ######################################
	// --- switch syntax

	switch 9 {
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("two or three")
	case 9:
		fmt.Println("nine")
	default:
		fmt.Println("other")
	}

	// --- there is a possibilty with initialiser
	// switch i := 2+3; i {}

	// --- possibilty to evaluate inside case

	i := 17
	switch {
	case i < 10:
		fmt.Println("less than 10")
	case i >= 10:
		fmt.Println("more or equal to 10")
	}

	// --- ! breaks are implicit inside switch !
	// but we can still use them if logic needs them

}
