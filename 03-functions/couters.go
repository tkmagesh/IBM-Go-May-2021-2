package main

import "fmt"

func main() {
	increment, decrement := getCounter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}

func getCounter() (func() int, func() int) {
	var counter int = 0

	increment := func() int {
		counter++
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}
	return increment, decrement
}
