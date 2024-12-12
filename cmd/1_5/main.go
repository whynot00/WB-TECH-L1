package main

import (
	"time"
	"wb-tech-l1/cmd/1_5/reader"
	"wb-tech-l1/cmd/1_5/writter"
)

const timeDelay = 2

var dataset = []string{"aaa", "aasd", "adsasd", "s", "f", "s", "f", "asf", "sad"}

func main() {

	dataCh := make(chan interface{})

	writter := writter.NewWritter(dataCh)
	reader := reader.NewReader(dataCh)

	go writter.StartWrite(dataset)
	go reader.StartRead()

	// shutdown
	func() {
		defer close(dataCh)

		time.Sleep(time.Second * time.Duration(timeDelay))

		writter.Stop()
	}()

}
