package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

const timeout = time.Second * 3

// Первый способ
func firstWay(done chan struct{}) {

	for {
		select {
		case <-done:
			fmt.Println("Горутина завершена.")
			return
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(time.Second)
		}
	}
}

// Второй способ
func secondWay(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Гортина завершена.")
			return
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(time.Second)
		}
	}
}

// Третий способ
func thirdWay(stopped *int32) {

	for atomic.LoadInt32(stopped) == 0 {
		fmt.Println("Горутина работает...")
		time.Sleep(time.Second)
	}

	fmt.Println("Горутина завершена.")

}

func main() {

	{
		fmt.Println("Первый способ запущен...")
		doneCh := make(chan struct{})

		go firstWay(doneCh)

		time.Sleep(timeout)
		doneCh <- struct{}{}

		// чтобы горутина успела завершить работу (лучше использовать waitgroup)
		time.Sleep(time.Second * 2)
	}

	{
		fmt.Println("Второй способ запущен...")
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		go secondWay(ctx)

		<-ctx.Done()

		// чтобы горутина успела завершить работу (лучше использовать waitgroup)
		time.Sleep(time.Second * 2)
	}

	{
		fmt.Println("Третий способ запущен...")

		var stopped int32
		go thirdWay(&stopped)

		time.Sleep(timeout)
		atomic.StoreInt32(&stopped, 1)

		// чтобы горутина успела завершить работу (лучше использовать waitgroup)
		time.Sleep(time.Second * 2)
	}
}
