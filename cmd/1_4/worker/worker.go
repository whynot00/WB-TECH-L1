package worker

import (
	"fmt"
)

const (
	StatusEnded = 1
)

type Worker struct {
	workerNum int

	dataCh    chan interface{}
	isRunning bool
}

func NewWorker(dataCh chan interface{}, workerNum int) *Worker {
	return &Worker{
		workerNum: workerNum,
		dataCh:    dataCh,
	}
}

func (w *Worker) Run() {
	// обработка входных данных через бесконечный цикл
	for !w.isRunning {
		data := <-w.dataCh
		fmt.Printf("workerNum: %d | data: %v\n", w.workerNum, data)
	}

}

func (w *Worker) Stop() {
	// прерываем бесконечный цикл
	w.isRunning = true
	fmt.Printf("worker %d was stopped\n", w.workerNum)
}
