package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{4, 1, 5, 2, 3}
	//var nos [5]int = [...]int{4, 1, 5, 2, 3}
	var nos = [...]int{4, 1, 5, 2, 3}
	fmt.Println(nos)
	fmt.Println(len(nos))

	//iteration
	/*
		for i := 0; i < len(nos); i++ {
			fmt.Println(nos[i])
		}
	*/

	//iteration using the range construct
	for idx, value := range nos {
		fmt.Printf("Value at index [%d] is %d\n", idx, value)
	}
}
