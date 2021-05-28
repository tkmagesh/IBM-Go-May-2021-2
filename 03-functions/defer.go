package main

import "fmt"

func main() {
	defer defferedMain()
	fn()
}

func defferedMain() {
	fmt.Println("deffered from main")
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println("Exiting after recovering")
	} else {
		fmt.Println("Exiting normally")
	}

}

func fn() {
	x := 200
	fmt.Println("fn started")
	defer func() {
		fmt.Println("deffered function (1) executed after fn", x)
	}()
	defer func() {
		fmt.Println("deffered function (2) executed after fn", x)
	}()

	defer func() {
		fmt.Println("deffered function (3) executed after fn", x)
	}()
	defer func() {
		fmt.Println("deffered function (4) executed after fn", x)
	}()
	f2()
	fmt.Println("fn completed")
}

func f2() {
	fmt.Println("f2 started")
	panic("Sorry! I cannot continue anymore!!")
}
