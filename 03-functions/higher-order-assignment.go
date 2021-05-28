package main

import "fmt"

func main() {
	x, y := 100, 200
	addResult := add(x, y)
	fmt.Println("addResult = ", addResult)
	subResult := subtract(x, y)
	fmt.Println("subResult = ", subResult)
	mulResult := multiply(x, y)
	fmt.Println("mulResult = ", mulResult)
	divResult := divide(x, y)
	fmt.Println("divResult = ", divResult)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

/*
DONOT modify the add, subtract, multiply, divide functions
DONOT repeat the logic for printing "processing 100 and 200"
desired output
processing 100 and 200
addResult =  300
processing 100 and 200
subResult =  -100
processing 100 and 200
mulResult =  20000
processing 100 and 200
divResult =  0

*/
