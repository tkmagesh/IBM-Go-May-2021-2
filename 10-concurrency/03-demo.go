package main

import (
	"fmt"
	"sync"
)

/* sharing memory for communication - NOT ADVICED */
var result int
var mutex sync.Mutex = sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Printf("Result = %d\n", result)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	result = x + y
	mutex.Unlock()
	wg.Done()
}
