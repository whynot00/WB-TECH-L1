package main

import "fmt"

func main() {
	fmt.Println(reverse("Глеба"))
}

func reverse(word string) string {

	runes := []rune(word)
	for leftRune, rightRune := 0, len(runes)-1; leftRune < rightRune; leftRune, rightRune = leftRune+1, rightRune-1 {
		runes[leftRune], runes[rightRune] = runes[rightRune], runes[leftRune]
	}

	return string(runes)
}
