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
	err2 := ioutil.WriteFile("data1-copy.txt", fileContents, 0x777)
	if err2 != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("File copied")
}
