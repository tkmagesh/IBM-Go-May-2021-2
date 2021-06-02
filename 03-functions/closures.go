package main

import (
	"fmt"
	"time"
)

func main() {
	for no := 1; no <= 10; no++ {
		go func(x int) {
			fmt.Println(x)
			time.Sleep(500 * time.Millisecond)
		}(no)
	}
	var input string
	fmt.Scanln(&input)
}
