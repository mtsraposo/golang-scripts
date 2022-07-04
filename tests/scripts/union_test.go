package tests

import (
	. "github.com/mtsraposo/xp-strats/scripts"
	. "github.com/mtsraposo/xp-strats/utils"
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		intervals [][]float64
		union     [][]float64
	}{
		{[][]float64{{2, 5}, {1, 3}, {7, 10}, {10, 15}, {-3.5, -1.5}, {20, 22}, {12, 20}},
			[][]float64{{-3.5, -1.5}, {1, 5}, {7, 22}}},
		{[][]float64{{2, 5}, {1, 1}, {1, 2}, {5, 5}, {6, 6}, {5, 6}},
			[][]float64{{1, 6}}},
	}
	for _, test := range tests {
		union := Union(test.intervals)
		if !MatricesEqual(union, test.union) {
			t.Errorf("Failed to mergeSort intervals %v. %v does not equal %v", test.intervals, union, test.union)
		}
		t.Logf("Union(%v) = %v", test.intervals, union)
	}
}
