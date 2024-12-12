package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	inStr := "abcd"

	fmt.Println(isUnique(inStr))
}

func isUnique(input string) bool {
	inRune := []rune(strings.ToLower(input))
	tempRune := make([]rune, 0, len(inRune))

	for _, r := range inRune {
		if slices.Contains(tempRune, r) {
			return false
		}
		tempRune = append(tempRune, r)
	}

	return true
}
