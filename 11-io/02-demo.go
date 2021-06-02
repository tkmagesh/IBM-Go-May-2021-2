package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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
	count, err := alphaReader.src.Read(p)
	fmt.Println(count)
	if err != nil {
		return count, err
	}

	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			continue
		} else {
			p[i] = 0
		}
	}
	return count, io.EOF
}

func main() {
	strReader := strings.NewReader("Hello! How are you?")
	alphaReader := NewAlphaReader(strReader)
	io.Copy(os.Stdout, alphaReader)
	fmt.Println()
}
