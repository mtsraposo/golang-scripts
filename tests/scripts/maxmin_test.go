package tests

import (
	. "github.com/mtsraposo/xp-strats/scripts"
	"testing"
)

func TestMaxMinBreaches(t *testing.T) {
	tests := []struct {
		prices         []float64
		maxMinBreaches [2]int
	}{
		{[]float64{20.50, 21.20, 22.40, 20.90, 21.57}, [2]int{2, 0}},
		{[]float64{20.52, 20.52, 19.81, 20.52}, [2]int{0, 1}},
		{[]float64{30.00, 30.12, 29.55, 28.43, 27.93, 31.15, 31.16, 30.98, 32.55}, [2]int{4, 3}},
	}

	for _, test := range tests {
		maxMinBreaches := CalcMaxMinBreaches(test.prices)
		if maxMinBreaches[0] != test.maxMinBreaches[0] || maxMinBreaches[1] != test.maxMinBreaches[1] {
			t.Errorf("failed to calculate maxMinBreaches for prices %v. %v did not equal %v",
				test.prices, maxMinBreaches, test.maxMinBreaches)
		}
		t.Logf("CalcMaxMinBreaches(%v) = %v", test.prices, maxMinBreaches)
	}
}
