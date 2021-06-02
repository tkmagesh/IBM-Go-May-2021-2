package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("data1.txt")
	if fileError != nil {
		log.Fatalln(fileError)
	}
	//defer fileHandle.Close()
	inputReader := bufio.NewReader(fileHandle)
	lineNo := 0
	for {
		line, err := inputReader.ReadString('\n')
		lineNo++
		fmt.Printf("Line # : %d, contents : %s\n", lineNo, line)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
