package main

import (
	"fmt"
	"sync"
)

var numbers = []int{2, 4, 6, 8, 10}

func main() {

	var (
		numbers1 []int
		numbers2 []int
	)

	numbers1 = append(numbers1, numbers...)

	// исполнение без использования каналов
	sumNum1 := squareWithoutChannels(numbers1)

	numbers2 = append(numbers2, numbers...)

	// исполнение с использованием каналов
	sumNum2 := squareWithChannels(numbers2)

	fmt.Println(sumNum1 == sumNum2)
}

func squareWithoutChannels(numbers []int) int {
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for index, num := range numbers {
		go func() {
			defer wg.Done()
			numbers[index] = num * num
		}()
	}

	wg.Wait()

	var sum int

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func squareWithChannels(numbers []int) int {
	sumCh := make(chan int) // Создаем канал для передачи результатов

	for _, num := range numbers {
		go func() {
			sumCh <- num * num // Отправляем квадрат числа в канал
		}()
	}

	sum := 0

	for i := 0; i < len(numbers); i++ { // Получаем результаты из канала
		sum += <-sumCh // Принимаем результат из канала и добавляем к сумме
	}

	return sum
}
