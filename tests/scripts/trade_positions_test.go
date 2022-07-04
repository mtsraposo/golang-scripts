package tests

import (
	. "github.com/mtsraposo/golang-scripts/scripts"
	. "github.com/mtsraposo/golang-scripts/utils"
	"testing"
	"time"
)

func TestTradePositions(t *testing.T) {
	day := 24 * time.Hour
	refDate := time.Date(2022, 7, 6, 0, 0, 0, 0, time.UTC)
	refTrades := []Trade{
		{refDate, "instr-1", true, 1, 1},
		{refDate, "instr-2", true, 1, 1},
		{refDate.Add(day), "instr-1", true, 2, 1},
		{refDate.Add(day), "instr-2", false, 1, 1},
		{refDate.Add(2 * day), "instr-1", false, 3, 1},
		{refDate.Add(2 * day), "instr-2", true, 1, 1},
	}
	tests := []struct {
		trades    []Trade
		date      time.Time
		positions []Position
	}{
		{refTrades, refDate,
			[]Position{
				{"instr-1", 1},
				{"instr-2", 1},
			},
		},
		{refTrades, refDate.Add(day),
			[]Position{
				{"instr-1", 3},
			},
		},
		{refTrades, refDate.Add(day * 2),
			[]Position{
				{"instr-2", 1},
			},
		},
	}
	for _, test := range tests {
		positions := TradePositions(test.trades, test.date)
		if !PositionsEqual(positions, test.positions) {
			t.Errorf("Failed to calculate positions for %v. %v does not equal %v",
				test.trades, positions, test.positions)
		} else {
			t.Logf("TradePositions(%v,%s) = %v", FormatTrades(test.trades), FormatDate(test.date), positions)
		}
	}

}
