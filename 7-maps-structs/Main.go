package main

import (
	"fmt"
	"reflect" // to work with field tags
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

	// --- ! the order of return of the map is not guaranteed !
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

	// --- naming conventions - scope
	// struct name:
	//   uppercase - export globally
	//   lowercase - package scope
	// field names:
	//   lowercase - package scope
	//   upppercase - global scope

	myself := Human{
		name:   "Antek",
		age:    21,
		height: 1.77,
		alive:  true,
	}

	fmt.Println(myself)

	// --- anonymous struct
	// myslef := struct{name string}{name: "antek"}

	// --- struct is copied by value

	// --- embedding
	// there is no inheritance in golang
	// you can put one type into another (embed)

	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal // embedded field
		Speed  float32
		CanFly bool
	}

	// --- how to define embedded field
	// method 1
	ptak := Bird{}
	ptak.Name = "koliber"
	ptak.Origin = "america"
	ptak.CanFly = true
	ptak.Speed = 100.2
	// method 2

	//ptak :=  Bird{
	//Animal : Animal{Name:"chicken", Origin: "poland ;p"}
	//Speed: 2.2
	//CanFly: false
	//}
	fmt.Println(ptak.Name)

	// --- tag of the field

	type Order struct {
		Price int `required_min:"100"` // set a tag, later we can access it with reflect
		Date  int
	}

	// access tag of a field
	t := reflect.TypeOf(Order{})
	field, _ := t.FieldByName("Price")

	fmt.Println(field.Tag)

}
