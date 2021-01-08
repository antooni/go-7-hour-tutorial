package main

import (
	"fmt"
)

func main() {
	// ###################################
	// --- declare map

	populacjaMiast := make(map[string]int)

	populacjaMiast = map[string]int{
		"Warszawa":     1790658,
		"Kraków":       779115,
		"Łódź":         679941,
		"Wrocław":      642869,
		"Poznań":       534813,
		"Gdańsk":       470907,
		"Zduńska Wola": 41686,
	}

	// --- the order of return of the map is not guaranteed
	// unlike in arrays

	// --- add to map
	populacjaMiast["Radom"] = 66666
	fmt.Printf("%v\n", populacjaMiast)

	// --- delete from map
	delete(populacjaMiast, "Radom")
	fmt.Printf("%v\n", populacjaMiast)

	// --- solve the problem of returning 0 where there is no matching key

	pop, ok := populacjaMiast["Radom"]
	// when there is no matching key
	// second variable will contain false
	fmt.Println(pop, ok)

	// --- length of map
	n := len(populacjaMiast)
	fmt.Println("len = ", n)

	// ##########################################
	// declare struct

	type Human struct {
		name   string
		age    int
		height float32
		alive  bool
	}

	myself := Human{
		name:   "Antek",
		age:    21,
		height: 1.77,
		alive:  true,
	}

	fmt.Println(myself)
}
