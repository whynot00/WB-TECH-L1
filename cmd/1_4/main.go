package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"wb-tech-l1/cmd/1_4/flags"
	"wb-tech-l1/cmd/1_4/worker"
)

// тестовые данные
var dataset = []string{"11", "asd", "gasdga", "kdmdgr", "asfkmag", "asfa"}

func main() {
	// изменить количество воркеров, запуск с флагом -wc {workersCount}
	// по умолчанию 5
	workersCount := flags.ParseFlag()

	isRunning := true
	workers := make([]*worker.Worker, workersCount)

	// канал для считывания сигналов
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)

	// канал из которого буду браться данные
	dataCh := make(chan interface{})

	// создаем необходимое количество воркеров
	for i := 0; i < workersCount; i++ {
		workers[i] = worker.NewWorker(dataCh, i)
	}

	// запускаем каждый воркер
	for _, worker := range workers {
		go worker.Run()
	}

	// засовываем тестовые данные
	go func() {
		count := 0
		for isRunning {
			count = runPool(dataset, dataCh, count)
		}
	}()

	// ждем сигнал
	<-stopCh
	fmt.Println()
	func() {
		// в конце закрываем канал
		defer close(dataCh)

		// прекращаем засовывать данные
		isRunning = false

		// останавливаем воркеры
		for _, worker := range workers {
			go worker.Stop()
		}

		// ждем время чтобы каждый воркер завершил свою работу
		time.Sleep(time.Second * 5)
	}()

}

func runPool(data []string, dataCh chan interface{}, count int) int {
	if count < len(data) {
		dataCh <- data[count]
	}
	return count + 1
}
