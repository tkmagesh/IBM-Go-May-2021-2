package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(fileContents))
}
