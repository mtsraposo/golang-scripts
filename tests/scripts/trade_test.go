package tests

import (
	"github.com/go-gota/gota/dataframe"
	. "github.com/mtsraposo/golang-scripts/scripts"
	"testing"
)

func TestTradeTotals(t *testing.T) {
	tests := []struct {
		trades    []Trade
		netAmount TradeDataFrame
	}{
		{
			trades: []Trade{
				{"2022-07-06", "instr-1", true, 1, 1},
				{"2022-07-06", "instr-2", false, 1, 1},
				{"2022-07-07", "instr-1", true, 2, 1},
				{"2022-07-07", "instr-1", false, 1, 1},
			},
			netAmount: TradeDataFrame{
				DataFrame: dataframe.LoadStructs(
					[]Total{
						{"2022-07-06", 0},
						{"2022-07-07", -1},
					}),
			},
		},
	}
	for _, test := range tests {
		net := TradeSummary(test.trades)
		if !DataFramesEqual(net, test.netAmount) {
			t.Errorf("Failed to calculate netAmount trade amounts for trades %v. %v does not equal %v",
				test.trades, net, test.netAmount)
		} else {
			t.Logf("TradeSummary(%v) = %v", test.trades, net)
		}
	}
}

func TestTradePositions(t *testing.T) {
	refTrades := []Trade{
		{"2022-07-06", "instr-1", true, 1, 1},
		{"2022-07-06", "instr-1", true, 1, 1},
		{"2022-07-06", "instr-2", true, 1, 1},
		{"2022-07-07", "instr-1", true, 2, 1},
		{"2022-07-07", "instr-2", false, 1, 1},
		{"2022-07-08", "instr-1", false, 4, 1},
		{"2022-07-08", "instr-2", true, 1, 1},
	}
	tests := []struct {
		trades    []Trade
		date      string
		positions TradeDataFrame
	}{
		{trades: refTrades, date: "2022-07-06",
			positions: TradeDataFrame{DataFrame: dataframe.LoadStructs([]Position{
				{"instr-1", 2},
				{"instr-2", 1},
			})},
		},
		{trades: refTrades, date: "2022-07-07",
			positions: TradeDataFrame{DataFrame: dataframe.LoadStructs([]Position{
				{"instr-1", 4},
			})},
		},
		{trades: refTrades, date: "2022-07-08",
			positions: TradeDataFrame{DataFrame: dataframe.LoadStructs([]Position{
				{"instr-2", 1},
			})},
		},
	}
	for _, test := range tests {
		positions := TradePositions(test.trades, test.date)
		if !DataFramesEqual(positions, test.positions) {
			t.Errorf("Failed to calculate positions for %v. %v does not equal %v",
				test.trades, positions, test.positions)
		} else {
			t.Logf("TradePositions(%v,%s) = %v", test.trades, test.date, positions)
		}
	}

}

func DataFramesEqual(calculated TradeDataFrame, expected TradeDataFrame) bool {
	if calculated.Nrow() != expected.Nrow() {
		return false
	}
	calculatedRecords := calculated.Records()
	for i, row := range expected.Records() {
		for j, elem := range row {
			if elem != calculatedRecords[i][j] {
				return false
			}
		}
	}
	return true
}