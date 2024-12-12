package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	inputStr := "snow dog sun"

	fmt.Printf("%s - %s", inputStr, reverseString(inputStr))
}

func reverseString(str string) string {
	tempSlice := strings.Split(str, " ")

	slices.Reverse(tempSlice)

	return strings.Join(tempSlice, " ")
}
