package main

import "fmt"

func xorBit(n int64, pos uint) int64 {

	mask_conjunction := ^(int64(1 << pos))
	mask_disjunction := int64(1 << pos)

	res := n | mask_disjunction
	if res != n {
		return res
	} else {
		return n & mask_conjunction
	}

}

func main() {
	num := int64(16)
	pos := uint(0)

	num = xorBit(int64(num), pos)

	fmt.Println(num)
}
