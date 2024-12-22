package main

import (
	"fmt"
	"math/rand"
)

// Функция для обмена двух элементов массива
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Функция для разбиения массива и выбора опорного элемента
func partition(arr []int, low int, high int) int {
	pivotIndex := rand.Intn(high - low + 1) // Случайный индекс опорного элемента
	pivotIndex += low                       // Приводим к диапазону [low, high]
	pivot := arr[pivotIndex]                // Опорный элемент
	swap(arr, pivotIndex, high)             // Перемещаем опорный элемент в конец

	i := low - 1 // Индекс меньшего элемента

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot { // Если текущий элемент меньше опорного
			i++
			swap(arr, i, j) // Меняем местами элементы
		}
	}

	swap(arr, i+1, high) // Помещаем опорный элемент на правильное место
	return i + 1         // Возвращаем индекс опорного элемента
}

// Рекурсивная функция QuickSort
func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high) // Разбиваем массив вокруг опорного элемента

		quickSort(arr, low, pi-1)  // Сортируем левую часть
		quickSort(arr, pi+1, high) // Сортируем правую часть
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Отсортированный массив:", arr)
}
