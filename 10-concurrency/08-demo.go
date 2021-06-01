/* buffered channels */

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	go fn(ch)
	fmt.Println("[main] Before reading data from channel")
	fmt.Println("[main] data from channel ", <-ch)
	fmt.Println("[main] After reading data from the channel")
	//time.Sleep(2 * time.Second)
}

func fn(ch chan int) {
	fmt.Println("[fn] Before writing 10 into the channel")
	ch <- 10
	fmt.Println("[fn] After writing 10 into the channel")
	fmt.Println("[fn] Before writing 20 into the channel")
	ch <- 20
	fmt.Println("[fn] After writing 20 into the channel")
	fmt.Println("[fn] Before writing 30 into the channel")
	ch <- 30
	fmt.Println("[fn] After writing 30 into the channel")
	fmt.Println("[fn] Before writing 40 into the channel")
	ch <- 40
	fmt.Println("[fn] After writing 40 into the channel")
}
