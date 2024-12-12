package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 14

	fmt.Println(binarySearch(numbers, target))
}

func binarySearch(array []int, target int) int {
	return binarySearchHidden(array, 0, len(array)-1, target)
}

func binarySearchHidden(array []int, startIndex, endIndex, target int) int {
	middleElem := (startIndex + endIndex) / 2
	if startIndex >= endIndex {
		return -1
	}
	if array[middleElem] == target {
		return middleElem
	}
	if array[middleElem] > target {
		return binarySearchHidden(array, startIndex, middleElem, target)
	}
	if array[middleElem] < target {
		return binarySearchHidden(array, middleElem+1, endIndex, target)
	}
	return -1
}
