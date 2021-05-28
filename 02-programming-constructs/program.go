package main

import "fmt"

func main() {
	//if else construct
	/*
		var no int = 20
		if no%2 == 0 {
			fmt.Printf("No %d is an even number\n", no)
		} else {
			fmt.Printf("No %d of type %T is an odd number\n", no, no)
		}
	*/

	if no := 20; no%2 == 0 {
		fmt.Printf("No %d is an even number\n", no)
	} else {
		fmt.Printf("No %d of type %T is an odd number\n", no, no)
	}

	//for construct
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum of integers from 1 to 10 = %d\n", sum)

	//for as a 'while' construct
	newSum := 1
	for newSum < 100 {
		newSum += newSum
	}
	fmt.Println(newSum)

	//indefinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Println(x)
}
