package main

import "fmt"

func sum(nos []int, ch chan int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	ch <- result
}

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	/* break the above set into two sets and add them concurrently and print the final result */
	firstSet := data[0 : len(data)/2]
	secondSet := data[len(data)/2:]
	ch1 := make(chan int)
	ch2 := make(chan int)
	go sum(firstSet, ch1)
	go sum(secondSet, ch2)
	result := <-ch1 + <-ch2
	fmt.Println(result)
}
