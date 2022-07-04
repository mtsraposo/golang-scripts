package tests

import (
	. "github.com/mtsraposo/golang-scripts/scripts"
	. "github.com/mtsraposo/golang-scripts/utils"
	"testing"
)

func TestCycles(t *testing.T) {
	tests := []struct {
		set    []int
		cycles [][]int
	}{
		{[]int{1, 2, 3, 4}, [][]int{{1}, {2}, {3}, {4}}},
		{[]int{4, 3, 2, 1}, [][]int{{1, 4}, {2, 3}}},
		{[]int{5, 1, 3, 4, 2}, [][]int{{1, 2, 5}, {3}, {4}}},
		{[]int{5, 3, 1, 2, 4}, [][]int{{1, 3, 2, 4, 5}}},
	}

	for _, test := range tests {
		cycles := PermutationCycles(test.set)
		if !MatricesEqual(cycles, test.cycles) {
			t.Errorf("failed to identify permutation cycles in %v. %v does not equal %v",
				test.set, cycles, test.cycles)
		} else {
			t.Logf("PermutationCycles(%v) = %v", test.set, cycles)
		}
	}
}
