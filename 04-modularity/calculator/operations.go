package calculator

import "fmt"

var operationCount int

func GetOperationCount() int {
	return operationCount
}

func log(msg string) {
	fmt.Println(msg)
}
