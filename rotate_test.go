package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tables := []struct {
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

	for _, table := range tables {
		rotated := rotate(table.L, table.k)
		if !arraysEqual(rotated, table.R) {
			t.Errorf("Rotation of (%v, %d) failed. %v does not equal %v.", table.L, table.k, rotated, table.R)
		}
	}
}

func arraysEqual(A []int, B []int) bool {
	if len(A) != len(B) {
		return false
	}

	for i, v := range A {
		if v != B[i] {
			return false
		}
	}

	return true
}