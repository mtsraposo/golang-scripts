package scripts

import (
	"sort"
	"time"
)

type Position struct {
	Instrument string
	Total      int
}

type DateToPosition map[time.Time]map[string]int

type DatePosition struct {
	Date      time.Time
	Positions map[string]int
}

func TradePositions(trades []Trade, date time.Time) []Position {
	dateToPositions := groupByInstrument(trades)
	positions := mapToArray(dateToPositions)
	sortByDate(positions)
	dateToPositions = accumulate(positions)
	positionsOnDate := getByDate(dateToPositions, date)
	return mapPositionsToArray(positionsOnDate)
}

func groupByInstrument(trades []Trade) DateToPosition {
	dateToPosition := make(DateToPosition)
	for _, trade := range trades {
		date := extractDate(trade)
		side := -1 * extractSideMultiplier(trade)
		if _, dateExists := dateToPosition[date]; dateExists {
			if _, positionExists := dateToPosition[date][trade.Instrument]; positionExists {
				dateToPosition[date][trade.Instrument] += trade.Quantity * side
			} else {
				dateToPosition[date][trade.Instrument] = trade.Quantity * side
			}
		} else {
			dateToPosition[date] = map[string]int{trade.Instrument: trade.Quantity * side}
		}
	}
	return dateToPosition
}

func mapToArray(toConvert DateToPosition) []DatePosition {
	var datePositions []DatePosition
	for date, positions := range toConvert {
		datePositions = append(datePositions, DatePosition{date, positions})
	}
	return datePositions
}

func sortByDate(positions []DatePosition) {
	sort.SliceStable(positions, func(i, j int) bool {
		return positions[i].Date.Before(positions[j].Date)
	})
}

func accumulate(positions []DatePosition) DateToPosition {
	cumulativePositions := make(DateToPosition)
	cumulativePositions[positions[0].Date] = positions[0].Positions

	previous := positions[0].Positions
	for _, date := range positions[1:] {
		current := getCurrentSum(date.Positions, previous)
		cumulativePositions[date.Date] = current
		previous = current
		current = map[string]int{}
	}
	return cumulativePositions
}

func getCurrentSum(positions map[string]int, previous map[string]int) map[string]int {
	current := make(map[string]int)
	for position, total := range positions {
		current[position] = total
		previousPos, previousExists := previous[position]
		if previousExists {
			current[position] += previousPos
		}
		if current[position] == 0 {
			delete(current, position)
		}
	}
	return current
}

func getByDate(cumulativePositions DateToPosition, date time.Time) map[string]int {
	return cumulativePositions[date]
}

func mapPositionsToArray(positionsMap map[string]int) []Position {
	var positionsArray []Position
	for position, total := range positionsMap {
		positionsArray = append(positionsArray, Position{position, total})
	}
	sort.SliceStable(positionsArray, func(i, j int) bool {
		return positionsArray[i].Instrument < positionsArray[j].Instrument
	})

	return positionsArray
}
