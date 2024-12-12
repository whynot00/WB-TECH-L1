package main

import (
	"fmt"
	"time"
)

const (
	millisecond = 1       // ms
	second      = 1000    // ms
	minute      = 60000   //ms
	hour        = 3600000 //ms

	// итд
)

func main() {

	startTime := time.Now()
	mySleep(second * 5)

	fmt.Println(time.Since(startTime))

}

func mySleep(timeDuration int64) {
	endTime := time.Now().UnixMilli() + timeDuration

	for {
		if endTime < time.Now().UnixMilli() {
			break
		}
	}
}
