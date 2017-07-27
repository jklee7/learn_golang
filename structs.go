package main

import (
	"fmt"
)

// STRUCTS
// Go doesn't have classes - instead it has structs.
// Structs are defined in the following manner:
// type [struct name] struct {}

type car struct {
	brand string
	model string
	color string
	doors int
}

// STRUCT CONSTRUCTORS
// Go doesn't have constructors in the typical OOP sense, but constructor
// functions are easy to implement.
// A constructor function creates and returns a pointer to a struct
func NewCar(brand string, model string, color string, doors int) *car {
	c := new(car)
	c.brand = brand
	c.model = model
	c.color = color
	c.doors = doors
	return c
}

func main() {
	// Structs can then be instantiated as you would any other variables
	hyundai_excel := car{}
	hyundai_excel.brand = "Hyundai"
	hyundai_excel.model = "Excel"
	hyundai_excel.color = "Red"
	hyundai_excel.doors = 2

	fmt.Println(hyundai_excel)

	// shortform method
	ford_falcon := car{"Ford", "Falcon", "Blue", 4}
	fmt.Println(ford_falcon)

	// You can also instantiate structs as pointers
	holden_commodore := new(car)
	(*holden_commodore).brand = "Holden"
	(*holden_commodore).model = "Commodore"
	(*holden_commodore).color = "Black"
	(*holden_commodore).doors = 4

	fmt.Println(*holden_commodore)

	// Instantiate a struct using a constructor function
	mazda_crx := *NewCar("Mazda", "CRX", "Green", 4)
	fmt.Println(mazda_crx)
}
