package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{4, 1, 5, 2, 3}
	//var nos [5]int = [...]int{4, 1, 5, 2, 3}
	var nos = [...]int{4, 1, 5, 2, 3}
	fmt.Println(nos)
	fmt.Println(len(nos))

	//iteration
	/*
		for i := 0; i < len(nos); i++ {
			fmt.Println(nos[i])
		}
	*/

	//iteration using the range construct
	for idx, value := range nos {
		fmt.Printf("Value at index [%d] is %d\n", idx, value)
	}

	//slice
	var names []string = []string{"Magesh", "Ganesh", "Suresh"}
	fmt.Println(names)
	fmt.Println("No of names = ", len(names))

	//adding new elements
	/*
		names = append(names, "Ramesh", "Rajesh")
		fmt.Println("Names after adding a new one => ", names)
	*/
	additionalNames := []string{"Ramesh", "Rajesh"}
	names = append(names, additionalNames...)
	fmt.Println("Names after adding additional names => ", names)

	//slicing
	/*
		[lo : hi] => from low to hi-1
		[:hi] => from 0 to hi-1
		[lo :] => from lo to len(list)-1 (last element)
		[lo : lo] => empty slize
		[lo : lo + 1] => element at 'lo'
		[:] => create a copy of the slice
	*/
	fmt.Println("names[1:4] => ", names[1:4])
	fmt.Println("names[:3] => ", names[:3])
	fmt.Println("names[3:] => ", names[3:])
	fmt.Println("names[2:2] => ", names[2:2])
	fmt.Println("names[2:3] => ", names[2:3])

	//removing an item
	names = names[1:]
	fmt.Println("After removing the first item => ", names)

	newNames := names
	newNames = append(newNames, "Philip")
	fmt.Println("newNames => ", newNames)
	fmt.Println("names => ", names)

	newNames[0] = "Sanjay"
	fmt.Println(names[0])

	fmt.Println("5 random numbers")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

}
