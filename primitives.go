/* A crash course on Go primitives */

package main

import (
	"fmt"
)

func main() {
	// VARIABLES
	// Variables are declared in the following manner:
	// var [variable name] [variable type]
	var myInteger int = 10
	var myFloat32 float32 = 11.0
	var myString string = "This is a string"

	println(myInteger, myFloat32, myString)

	// Go also has a shortform method of variable declaration -
	// [variable name] := [value]
	// The variable type is implicitly assigned based on what value is assigned
	newInt := 10
	newFloat32 := 11.0
	newString := "This is a new string"

	println(newInt, newFloat32, newString)

	// CONSTANTS
	// Constants are declared in the following manner:
	// const [variable name] = [value]
	// You can declare multiple constants in one go like this:
	const (
		first  = 1
		second = 2
	)

	// or declare a single constant:
	const third = 3
	println(first, second, third)

	// IOTAS
	// Iotas are used in const declarations to simplify definitions of incrementing numbers.
	// The values of iota is reset to 0 whenever the reserved word const appears in the source.
	const (
		iota0 = iota //0
		iota1        //1
		iota2        //2
	)

	println(iota0, iota1, iota2)

	// Arithmetic can also be used to affect the values assigned to iotas
	const (
		iota_a = iota * 5 // 0x5
		iota_b            //1x5
		iota_c            //2x5
	)

	println(iota_a, iota_b, iota_c)

	// COMPLEX NUMBERS
	// Complex numbers have real and imaginary parts
	var complexNum complex64
	complexNum = complex(1, 2)
	fmt.Println(real(complexNum), imag(complexNum))
}
