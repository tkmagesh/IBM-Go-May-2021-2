package main

import (
	"fmt"
	"time"
)

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	go print("Hello")
	go print("World")
	/*
		var input string
		fmt.Println("press ENTER to exit")
		fmt.Scanln(&input)
	*/
	time.Sleep(2 * time.Second)
}
