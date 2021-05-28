package main

import "fmt"

func main() {
	/*
		fn := func() {
			fmt.Println("fn is invoked")
		}
		fn()
	*/
	/*
		fn := myFn
		fn()
	*/

	var fn func() = myFn
	fn()

	/* var addFn func(int, int) int = add
	fmt.Println(addFn(100, 200)) */

	func() {
		fmt.Println("fn is invoked")
	}()

	addFn := getAdder()
	fmt.Println(addFn(100, 200))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(400))

	printCalculatedResult(10, 20, add)

}

func myFn() {
	fmt.Println("fn is invoked")
}

func add(x, y int) (result int) {
	result = x + y
	return
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func printCalculatedResult(x, y int, fn func(int, int) int) {
	fmt.Printf("Result of processing %d and %d is %d\n", x, y, fn(x, y))
}
