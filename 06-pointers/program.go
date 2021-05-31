package main

import "fmt"

func main() {
	no := 10
	var noPtr *int = &no
	fmt.Printf("Address of %d is %v\n", no, noPtr)
	fmt.Printf("Value at address %v is %d\n", noPtr, *noPtr)
	/* no == *(&no) */

	no1, no2 := 10, 20
	fmt.Printf("Before swapping, no1 = %d  and no2 = %d\n", no1, no2)
	swap(&no1, &no2)
	fmt.Printf("After swapping, no1 = %d  and no2 = %d\n", no1, no2)
}

func swap(ptr1, ptr2 *int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}
