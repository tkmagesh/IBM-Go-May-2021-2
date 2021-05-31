package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//generate random numbers
	randomNos := generateRandomNos()
	fmt.Println("Random nos => ", randomNos)

	//find out the even numbers
	//evenNos := filterEvenNos(randomNos)
	evenNos := filter(randomNos, isEven)
	fmt.Println("Even Nos => ", evenNos)

	//find out the prime numbers
	//primeNos := filterPrimeNos(randomNos)
	primeNos := filter(randomNos, isPrime)
	fmt.Println("Prime Nos => ", primeNos)
}

func generateRandomNos() []int {
	randomNos := [50]int{}
	for i := 0; i < 50; i++ {
		randomNos[i] = rand.Intn(100)
	}
	return randomNos[:]
}

func filter(randomNos []int, criteria func(int) bool) []int {
	result := []int{}
	for _, randomNo := range randomNos {
		if criteria(randomNo) {
			result = append(result, randomNo)
		}
	}
	return result
}

func isEven(no int) bool {
	return no%2 == 0
}

func isPrime(no int) bool {
	if no <= 2 {
		return true
	} else {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				return false
			}
		}
		return true
	}
}

/*
func filterEvenNos(randomNos []int) []int {
	result := []int{}
	for _, randomNo := range randomNos {
		if randomNo%2 == 0 {
			result = append(result, randomNo)
		}
	}
	return result
}

func filterPrimeNos(randomNos []int) []int {
	result := []int{}
	for _, randomNo := range randomNos {
		if randomNo <= 2 {
			result = append(result, randomNo)
		} else {
			isPrime := true
			for i := 2; i <= (randomNo / 2); i++ {
				if randomNo%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				result = append(result, randomNo)
			}
		}
	}
	return result
}
*/
