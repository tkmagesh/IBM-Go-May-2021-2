package calculator

func Add(x, y int) int {
	operationCount++
	log("Add operation being carried out")
	return x + y
}
