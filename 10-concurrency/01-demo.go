package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	go print("Hello")
	go print("World")
}
