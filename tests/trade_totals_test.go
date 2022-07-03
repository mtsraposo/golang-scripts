package tests

import (
	"testing"
	"time"
)

func TestTradeTotals(t *testing.T) {
	refDate := time.Date(2022, 7, 4, 0, 0, 0, 0, time.UTC)
	day := 24 * time.Hour
	tests := []struct {
		trades []Trade
		net    []Net
	}{
		{[]Trade{{refDate, "instr-1", true, 1, 1},
			{refDate, "instr-2", false, 1, 1},
			{refDate.Add(day), "instr-1", true, 2, 1},
			{refDate.Add(day), "instr-1", false, 1, 1}},
			[]Net{{refDate, 0}, {refDate.Add(day), -1}}},
	}
	for _, test := range tests {
		net := tradeSummary(test.trades)
		if !totalsAreEqual(net, test.net) {
			t.Errorf("Failed to calculate net trade amounts for trades %v. %v does not equal %v",
				test.trades, net, test.net)
		}
	}
}

func totalsAreEqual(n1 []Net, n2 []Net) bool {
	if len(n1) != len(n2) {
		return false
	}
	for i, net := range n2 {
		if net.Date != n2[i].Date || net.Total != n2[i].Total {
			return false
		}
	}
	return true
}
