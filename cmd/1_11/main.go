package main

import (
	"wb-tech-l1/cmd/1_11/set"
)

func main() {
	set1 := set.NewSet()
	set1.Add(123, 111, 4123, 4444, 1231, 123)

	set2 := set.NewSet()
	set2.Add(111, 123, 4555, 123)

	set3 := set2.Intersection(set1)

	set3.Show()
}
