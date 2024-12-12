package main

import "fmt"

func main() {

	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	gradMap := getMap(nums)

	fmt.Println(gradMap)
}

func getMap(nums []float64) map[int][]float64 {

	resMap := make(map[int][]float64)

	for _, num := range nums {
		key := int(num) / 10 * 10
		resMap[key] = append(resMap[key], num)
	}
	return resMap
}
