package main

import (
	"fmt"
	"io"
	"os"
)

/* Using the io.Reader  */

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(alphaReader); i++ {
		if (alphaReader[i] >= 'A' && alphaReader[i] <= 'Z') || (alphaReader[i] >= 'a' && alphaReader[i] <= 'z') {
			p[count] = alphaReader[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	alphaReader := AlphaReader("Hello! How are you?")
	io.Copy(os.Stdout, &alphaReader)
	fmt.Println()
}
