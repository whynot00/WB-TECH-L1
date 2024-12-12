package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	counter int64
	m       sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()

	c.counter++
}

func (c *Counter) Value() int64 {
	return c.counter
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for b := 0; b < 10; b++ {
				counter.Increment()
				time.Sleep(time.Second)
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.Value())
}
