package main

import (
	"fmt"
)

func multiplyByTwo(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}

func printResults(in <-chan int) {
	for result := range in {
		fmt.Println(result)
	}
}

func main() {
	inputs := []int{1, 2, 3, 4, 5}

	inChan := make(chan int)
	outChan := make(chan int)

	go func() {
		for _, num := range inputs {
			inChan <- num
		}
		close(inChan)
	}()

	go multiplyByTwo(inChan, outChan)
	printResults(outChan)
}
