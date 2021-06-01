package main

import (
	"fmt"
	"time"
)

func print(msg string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		time.Sleep(delay)
		fmt.Println(msg)
	}
}

func main() {
	print("Hello", 2*time.Second)
	print("World", 3*time.Second)
	fmt.Println("press ENTER to exit..")
	var input string
	fmt.Scanln(&input)
}
