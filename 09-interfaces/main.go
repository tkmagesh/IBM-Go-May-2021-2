package main

import "fmt"

type User struct {
	name string
}

func main() {
	var x interface{}
	x = 10
	x = "string"
	x = true
	fmt.Println(x)

	x = 100
	if no, ok := x.(int); ok {
		fmt.Println(no * 2)
	} else {
		fmt.Println("value cannot be converted to int")
	}

	//x = "Magesh"
	x = User{"Magesh"}
	switch val := x.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", val, 2*val)
	case string:
		fmt.Printf("Length of the string is %v\n", len(val))
	case User:
		fmt.Println("User name is ", val.name)
	default:
		fmt.Println("Unknown type")
	}

}
