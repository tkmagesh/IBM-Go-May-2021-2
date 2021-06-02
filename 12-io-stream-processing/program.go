package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {

	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)
	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)
	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenSumCh, processWg)
	go sum(oddCh, oddSumCh, processWg)
	go merger(evenSumCh, oddSumCh, "result.txt", processWg)
	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Done")
}

func source(filename string, out chan int, wg *sync.WaitGroup) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
		wg.Done()
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		out <- val
	}

}

func splitter(dataCh chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
	close(evenCh)
	close(oddCh)
}

func sum(in chan int, out chan int, wg *sync.WaitGroup) {
	result := 0
	for no := range in {
		result += no
	}
	out <- result
	wg.Done()
}

func merger(evenSumCh chan int, oddSumCh chan int, resultFile string, wg *sync.WaitGroup) {
	defer wg.Done()
	rf, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	for i := 0; i < 2; i++ {
		select {
		case oddResult := <-oddSumCh:
			rf.Write([]byte(fmt.Sprintf("Sum of odd numbers : %d\n", oddResult)))
		case evenResult := <-evenSumCh:
			rf.Write([]byte(fmt.Sprintf("Sum of even numbers : %d\n", evenResult)))
		}
	}
}
