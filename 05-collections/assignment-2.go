package main

import "fmt"

var operations = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	/*
		display menu and accept the menu choice
		if the choice exists,
			accept the operands
			perform the operation
			print the result
		else
			exit
	*/

	for {
		operationChoice := getUserChoice()
		if operation, exists := operations[operationChoice]; exists {
			no1, no2 := getOperands()
			result := operation(no1, no2)
			fmt.Println("Result = ", result)
		} else {
			return
		}
	}

}

func getOperands() (int, int) {
	no1, no2 := 0, 0
	fmt.Println("Enter the operands")
	fmt.Scanf("%d %d\n", &no1, &no2)
	return no1, no2
}

func getUserChoice() int {
	choice := 0
	fmt.Println("Select the operation : ")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanf("%d\n", &choice)
	return choice
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
