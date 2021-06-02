package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type channelWriter struct {
	Channel chan byte
}

func NewChannelWriter() *channelWriter {
	return &channelWriter{
		Channel: make(chan byte, 1024),
	}
}

func (c *channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	go func() {
		defer close(c.Channel) // when done
		for _, b := range p {
			c.Channel <- b
		}
	}()

	return len(p), nil
}

func main() {
	cw := NewChannelWriter()
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		io.Copy(cw, fileReader)
	}()

	for c := range cw.Channel {
		fmt.Printf("%c\n", c)
	}
}
