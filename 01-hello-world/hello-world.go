//package declaration
package main

import "fmt"

//import dependent packages

//package level variables

//package init function

//main function
func main() {

	/*
		var msg string
		msg = "Hello world!"
		fmt.Println(msg)
	*/

	/*
		var msg string = "Hello world!"
		fmt.Println(msg)
	*/
	/*
		var msg = "Hello world!"
		fmt.Println(msg)
	*/

	//the following declaration is valid only in a function
	/*
		msg := "Hello world!"
		fmt.Println(msg)
	*/

	//declaring more than one variable
	/*
		var name string
		var msg string
		name = "Magesh"
		msg = "Hi"
		fmt.Println(msg, name)
	*/

	/*
		var (
			name string
			msg  string
		)
		name = "Magesh"
		msg = "Hi"
		fmt.Println(msg, name)
	*/

	/*
		var (
			name, msg string
		)
		name = "Magesh"
		msg = "Hi"
		fmt.Println(msg, name)
	*/

	/*
		var name, msg string
		name = "Magesh"
		msg = "Hi"
		fmt.Println(msg, name)
	*/

	/*
		var name, msg string = "Magesh", "Hi"
		fmt.Println(msg, name)
	*/

	/*
		var name, msg = "Magesh", "Hi"
		fmt.Println(msg, name)
	*/

	name, msg := "Magesh", "Hi"
	//fmt.Println(msg, name)
	/*
		output := fmt.Sprintf("%s %s!", msg, name)
		fmt.Printf("The generated msg is %q \n", output)
	*/

	fmt.Printf("The generated msg is \"%s %s!\" \n", msg, name)
}

//other functions
