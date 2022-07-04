package utils

import (
	. "github.com/mtsraposo/golang-scripts/scripts"
	"time"
)

func PositionsEqual(p1 []Position, p2 []Position) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i, pos := range p2 {
		if pos.Instrument != p2[i].Instrument || pos.Total != p2[i].Total {
			return false
		}
	}
	return true
}

func ArraysEqual[V int | float64](A []V, B []V) bool {
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

func MatricesEqual[V int | float64](A [][]V, B [][]V) bool {
	if len(A) != len(B) {
		return false
	}
	for i, row := range B {
		if !ArraysEqual(row, A[i]) {
			return false
		}
	}
	return true
}

type FormattedTrade struct {
	TradeDate  string
	Instrument string
	SideIsBuy  bool
	Quantity   int
	Price      int
}

func FormatTrades(trades []Trade) []FormattedTrade {
	var fmtTrades []FormattedTrade
	var fmtTrade FormattedTrade
	for _, trade := range trades {
		fmtTrade = FormattedTrade{
			FormatDate(trade.TradeDate),
			trade.Instrument,
			trade.SideIsBuy,
			trade.Quantity,
			trade.Price,
		}
		fmtTrades = append(fmtTrades, fmtTrade)
	}
	return fmtTrades
}

const layoutISO = "2006-01-02"

func FormatDate(date time.Time) string {
	return date.Format(layoutISO)
}
