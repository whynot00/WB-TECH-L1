package main

import "wb-tech-l1/cmd/1_12/set"

var inputData = []string{"cat", "cat", "dog", "cat", "tree"}

func main() {

	set := set.NewSet()

	set.Add(inputData...)

	set.Show()
}
