package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	done := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-done:
			fmt.Println("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
