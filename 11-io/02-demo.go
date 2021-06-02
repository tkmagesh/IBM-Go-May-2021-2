package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/* Chaining readers  */

type AlphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *AlphaReader {
	return &AlphaReader{source}
}

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	inputData := make([]byte, len(p))
	count, err := alphaReader.src.Read(inputData)
	if err != nil {
		return count, err
	}
	dataCount := 0
	for i := 0; i < len(inputData); i++ {
		if (inputData[i] >= 'A' && inputData[i] <= 'Z') || (inputData[i] >= 'a' && inputData[i] <= 'z') {
			p[dataCount] = inputData[i]
			dataCount++
		}
	}
	return dataCount, io.EOF
}

func main() {
	//strReader := strings.NewReader("Hello! How are you?")
	//alphaReader := NewAlphaReader(strReader)
	fileReader, err := os.Open("./data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	alphaReader := NewAlphaReader(fileReader)
	io.Copy(os.Stdout, alphaReader)
	fmt.Println()
}
