package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
	"strconv"

	"dummyModule/dummy"

	"github.com/fatih/color"
)

func main() {
	color.Yellow(strconv.Itoa(calc.Add(100, 200)))
	color.Green(strconv.Itoa(calc.Subtract(100, 200)))
	color.Red(fmt.Sprintf("Operation count : %d\n", calc.GetOperationCount()))
	utils.SayHi()
	dummy.WhoAmI()
}
