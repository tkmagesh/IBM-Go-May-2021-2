package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go fibonacci(ch)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	var input string
	fmt.Println("press ENTER to exit")
	fmt.Scanln(&input)
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	//generate the fibonacci numbers
	x, y := 0, 1
	stop := time.After(10 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}

}
