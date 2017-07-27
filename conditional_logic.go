package main

func main() {
	// IF
	// The syntax for If statements in Go is pretty much what
	// you'd expect for a C-type langauge.
	someFlag := true

	// Other valid comparisons: >, >=, <, <=
	if someFlag == true {
		println("True")
	} else {
		println("False")
	}

	// SWITCH
	option := 5
	switch option {
	case 1:
		println("1")
	case 2:
		println("2")
	default:
		println("Not 1 or 2")
	}

}
