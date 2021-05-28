package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/utils"

	"dummyModule/dummy"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	color.Red(fmt.Sprintf("Operation count : %d\n", calc.GetOperationCount()))
	utils.SayHi()
	dummy.WhoAmI()
}
