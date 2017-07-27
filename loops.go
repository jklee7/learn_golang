package main

func main() {
	// FOR LOOPS
	// Go only has 1 type of loop - the For loop

	// The traditional For loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// Can also be used like a while loop
	count := 0
	for ; count < 5; count++ {
		println(count)
	}

	// Or like an infinite loop
	for {
		if count > 5 {
			break
		}
		println(count)
		count++
	}

	// Looping through collections
	array1 := [3]int{1, 2, 3}
	for i := range array1 {
		println(i)
	}

	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	for i := range map1 {
		println(i)
	}
}
