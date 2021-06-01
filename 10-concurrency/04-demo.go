package main

import (
	"fmt"
	"sync"
)

/* communicate by sharing memory (using channels) */
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Printf("Result = %d\n", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	//writing result to the channel
	ch <- result
	wg.Done()
}
