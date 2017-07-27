package main

import (
	"fmt"
)

func main() {
	// ARRAYS
	// Arrays have a fixed length, which can't be changed after declaration
	var array1 [3]int
	array1[0] = 0
	array1[1] = 1
	array1[2] = 2

	array2 := [...]int{3, 4, 5}

	fmt.Println(array1)
	fmt.Println(array2)

	// SLICES
	// Slices are arrays with a dynamic length
	/* The way Go implements slices is by creating an array of exactly
	the right size. Every time elements are added or removed, Go will have
	to copy the new value to a new array in the background, which is
	very inefficient when dealing with multiple adds or removes. */

	/* We can get around this inefficiency by creating arrays with a larger
	initial length */

	var slice1 []int
	slice1 = make([]int, 2)     //create slice with initial length of 2 - [0,0]
	slice1 = append(slice1, 10) //[0,0,10]
	fmt.Println(slice1)

	slice2 := []int{10, 11, 12} //shorthand method of declaring and initialising slices
	fmt.Println(slice2)

	// MAPS
	// Maps are a collection of key-value pairs
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2

	map2 := make(map[int]string) //shorthand method of declaring maps
	map2[1] = "one"
	map2[2] = "two"

	fmt.Println(map1)
	fmt.Println(map2)
}
