package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	stop := make(chan string)
	go fibonacci(ch, stop)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	var input string
	fmt.Println("press ENTER to stop")
	fmt.Scanln(&input)
	stop <- "done"
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func fibonacci(ch chan int, stop chan string) {
	//generate the fibonacci numbers
	x, y := 0, 1
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
