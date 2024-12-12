package main

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testTable := []struct {
		arr      []int
		target   int
		expected int
	}{
		{
			arr:      []int{1, 2, 4, 5, 6, 78, 99, 124, 5236, 75465, 75975},
			target:   78,
			expected: 5,
		},
		{
			arr:      []int{1, 2, 4, 5, 6, 7, 8, 9, 9, 12, 45, 236, 7546, 57592},
			target:   1,
			expected: 0,
		},
		{
			arr:      []int{1, 2, 4, 5, 6, 7, 8, 9, 9, 12, 45, 236, 7546, 57592},
			target:   11,
			expected: -1,
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			target:   13,
			expected: 12,
		},
		{
			arr:      []int{1425, 2523, 6234, 6234, 7452, 8129, 9125, 10415, 11519},
			target:   7452,
			expected: 4,
		},
	}

	for _, test := range testTable {
		out := binarySearch(test.arr, test.target)

		if !reflect.DeepEqual(out, test.expected) {
			t.Errorf("\nArr: %v\nTarget: %d, Expected: %d, Output: %d", test.arr, test.target, test.expected, out)
		}
	}

}
