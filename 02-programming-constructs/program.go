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

	//switch case
	var y int = 21
	switch y % 2 {
	case 0:
		fmt.Println("Even number")
	case 1:
		fmt.Println("Odd number")
	}

	/*
		Rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		else => "Unknown score"
	*/

	score := 7
	/*
		switch score {
		case 0:
		case 1:
		case 2:
		case 3:
			fmt.Println("Terrible")
		case 4:
		case 5:
		case 6:
		case 7:
			fmt.Println("Not bad!")
		case 8:
		case 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown score")

		}
	*/

	/*
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not bad!")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Unknown score")

		}
	*/

	switch {
	case score >= 0 && score <= 3:
		fmt.Println("Terrible")
	case score >= 4 && score <= 7:
		fmt.Println("Not bad!")
	case score >= 8 && score <= 9:
		fmt.Println("Good")
	case score == 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}

	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println(" is <= 1")
		fallthrough
	case 2:
		fmt.Println(" is <= 2")
		fallthrough
	case 3:
		fmt.Println(" is <= 3")
		fallthrough
	case 4:
		fmt.Println(" is <= 4")
		fallthrough
	case 5:
		fmt.Println(" is <= 5")
		fallthrough
	case 6:
		fmt.Println(" is <= 6")
		fallthrough
	case 7:
		fmt.Println(" is <= 7")
	}
}
