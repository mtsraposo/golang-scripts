package main

import (
	"testing"
)

func TestMaxMinBreaches(t *testing.T) {
	tests := []struct {
		prices         []float32
		maxMinBreaches [2]int
	}{
		{[]float32{20.50, 21.20, 22.40, 20.90, 21.57}, [2]int{2, 0}},
		{[]float32{20.52, 20.52, 19.81, 20.52}, [2]int{0, 1}},
		{[]float32{30.00, 30.12, 29.55, 28.43, 27.93, 31.15, 31.16, 30.98, 32.55}, [2]int{4, 3}},
	}

	for _, test := range tests {
		maxMinBreaches := calcMaxMinBreaches(test.prices)
		if maxMinBreaches[0] != test.maxMinBreaches[0] || maxMinBreaches[1] != test.maxMinBreaches[1] {
			t.Errorf("failed to calculate maxMinBreaches for prices %v. %v did not equal %v",
				test.prices, maxMinBreaches, test.maxMinBreaches)
		}
	}
}
