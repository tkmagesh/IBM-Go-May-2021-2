package main

import "fmt"

func main() {
	fmt.Println("Generic Sum fn")
	fmt.Println(sum(10, 20))         //=> 30
	fmt.Println(sum(10, 20, 30, 40)) //=> 100
	fmt.Println(sum(10))             //=> 10
	fmt.Println(sum())               //=> 0

	nos := []int{10, 20, 30, 40}
	fmt.Println(sum(nos...))
	/*
		fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
		fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
		fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
		fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
		fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
	*/
}

func sum(nos ...int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result
}
