package main

func main() {
	// ADDITION
	addNum := 1 + 1 //2
	addNum++        //3
	addNum += 7     //10
	println(addNum)

	// SUBTRACTION
	subtractNum := 5 - 1 //4
	subtractNum--        //3
	subtractNum -= 3     //0
	println(subtractNum)

	// MULTIPLICATION
	multiplyNum := 2 * 5 //10
	multiplyNum *= 5     //50
	println(multiplyNum)

	// DIVISION
	divideNum := 100 / 10 //10
	divideNum /= 2        //5
	println(divideNum)

	// STRING CONCATENATION
	stringConcat := "String1" + " " + "String2"
	println(stringConcat)
}
