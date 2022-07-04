package scripts

import (
	"sort"
	"time"
)

// Trade Using int for monetary calculations to avoid rounding errors
// Trade A scale of precision (e.g. 1e6) may be used for conversion.
type Trade struct {
	TradeDate  time.Time
	Instrument string
	SideIsBuy  bool
	Quantity   int
	Price      int
}

type Net struct {
	Date  time.Time
	Total int
}

func TradeSummary(trades []Trade) []Net {
	dateToTotal := mapDateToTotal(trades)
	return genTotals(dateToTotal)
}

func mapDateToTotal(trades []Trade) map[time.Time]int {
	dateToTotal := make(map[time.Time]int)
	var date time.Time
	var sideMult int
	for _, trade := range trades {
		date = extractDate(trade)
		sideMult = extractSideMultiplier(trade)
		dateToTotal[date] += trade.Quantity * trade.Price * sideMult
	}
	return dateToTotal
}

func extractDate(trade Trade) time.Time {
	tradeDate := trade.TradeDate
	return time.Date(tradeDate.Year(), tradeDate.Month(), tradeDate.Day(), 0, 0, 0, 0, time.UTC)
}

func extractSideMultiplier(trade Trade) int {
	if trade.SideIsBuy {
		return -1
	}
	return 1
}

func genTotals(dateToTotal map[time.Time]int) []Net {
	var totals []Net
	for date, total := range dateToTotal {
		totals = append(totals, Net{date, total})
	}

	sort.SliceStable(totals, func(i, j int) bool {
		return totals[i].Date.Before(totals[j].Date)
	})

	return totals
}
