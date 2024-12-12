package main

import "fmt"

func main() {

	array := []string{
		"aa", "ss", "dd", "qq",
	}

	array = delByIndex2(array, 0)
	fmt.Println(array)

	array = delByIndex1(array, 1)
	fmt.Println(array)
}

// важен порядок
func delByIndex1(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

// порядок не важен
func delByIndex2(array []string, index int) []string {
	array[index] = array[len(array)-1]
	return array[:len(array)-1]
}
