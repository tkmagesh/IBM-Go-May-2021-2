package main

import (
	"errors"
	"fmt"
)

func main() {
	//helloWorld()
	//helloWorld("Hello World!")
	fmt.Println(add(100, 200))
	//fmt.Println(divide(21, 5))

	/*
		quotient, remainder := divide(21, 5)
		fmt.Println(quotient, remainder)
	*/
	quotient, _, err := divide(21, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotient)
	}

}

/*
func helloWorld() {
	fmt.Println("Hello World!")
}
*/

/*
func helloWorld(msg string) {
	fmt.Println(msg)
}
*/

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = errors.New("Invalid argument. Cannot divide by zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
