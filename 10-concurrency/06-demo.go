package main

import (
	"fmt"
	"time"
)

func print(msg string, delay time.Duration, ch1 chan string, ch2 chan string) {
	for i := 0; i < 5; i++ {
		<-ch1
		time.Sleep(delay)
		fmt.Println(msg)
		ch2 <- "done"
	}
}

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	go print("Hello", 1*time.Second, ch1, ch2)
	go print("World", 1*time.Second, ch2, ch1)
	ch1 <- "Start"
	fmt.Println("press ENTER to exit..")
	var input string
	fmt.Scanln(&input)
}

/*
Desired Output:
Hello
World
Hello
World
...

*/
