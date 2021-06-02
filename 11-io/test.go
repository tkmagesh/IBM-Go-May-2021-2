package main

import "fmt"

type MyInt int

func (myInt MyInt) WhoAmI() string {
	return "I am MyInt"
}

func main() {
	x := MyInt(123)
	fmt.Println(x.WhoAmI())
}
