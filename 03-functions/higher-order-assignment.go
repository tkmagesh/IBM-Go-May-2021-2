package main

import "fmt"

func main() {
	loggedAdd := loggedOper("Add", add)
	loggedAdd(100, 200)
	loggedAdd(10, 20)

	loggedOper("Subtract", subtract)(100, 200)
	loggedOper("Multiply", multiply)(100, 200)
	loggedOper("Divide", divide)(100, 200)
}

func loggedOper(operName string, oper func(int, int) int) func(int, int) {
	return func(x, y int) {
		fmt.Printf("Proessing %d and %d\n", x, y)
		result := oper(x, y)
		fmt.Printf("%s Result = %d\n", operName, result)
	}
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
