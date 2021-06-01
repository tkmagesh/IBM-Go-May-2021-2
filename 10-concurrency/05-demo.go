package main

import (
	"fmt"
)

/* communicate by sharing memory (using channels) */
func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	go add(300, 400, ch)
	result1 := <-ch
	result2 := <-ch
	fmt.Println(result1, result2)
	fmt.Printf("Result = %d\n", result1+result2)
}

func add(x, y int, ch chan int) {
	result := x + y
	//writing result to the channel
	ch <- result
}
