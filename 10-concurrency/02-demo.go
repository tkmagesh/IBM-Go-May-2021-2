package main

import (
	"fmt"
	"sync"
	"time"
)

func print(msg string, duration time.Duration, wg *sync.WaitGroup) {
	time.Sleep(duration)
	fmt.Println(msg)
	wg.Done()
}

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg.Add(2)
	go print("Hello", 3*time.Second, wg)
	go print("World", 5*time.Second, wg)
	/*
		var input string
		fmt.Println("press ENTER to exit")
		fmt.Scanln(&input)
	*/
	/*
		time.Sleep(2 * time.Second)
	*/
	wg.Wait()
	fmt.Println("Exiting the program")
}
