package main

import (
	"fmt"
	"os"
	"sync"
)

var numbers = []int{2, 4, 6, 8, 10}

func main() {
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, num := range numbers {
		go func() {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, num*num)
		}()
	}

	wg.Wait()
}
