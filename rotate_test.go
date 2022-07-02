package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		L []int
		k int
		R []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, -2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, -3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, -6, []int{5, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		rotated := rotate(test.L, test.k)
		if !arraysEqual(rotated, test.R) {
			t.Errorf("Rotation of (%v, %d) failed. %v does not equal %v.", test.L, test.k, rotated, test.R)
		}
	}
}