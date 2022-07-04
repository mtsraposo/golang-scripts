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
	datePositions := mapToArray(dateToPositions)
	sortByDate(datePositions)
	dateToPositionsCumulative := accumulate(datePositions)
	cumulativePositionsOnDate := getByDate(dateToPositionsCumulative, date)
	return mapSingleDatePositionsToArray(cumulativePositionsOnDate)
}

func groupByInstrument(trades []Trade) DateToPosition {
	dateToPositions := make(DateToPosition)
	for _, trade := range trades {
		date := extractDate(trade)
		side := -1 * extractSideMultiplier(trade)
		if _, dateExists := dateToPositions[date]; dateExists {
			if _, positionExists := dateToPositions[date][trade.Instrument]; positionExists {
				dateToPositions[date][trade.Instrument] += trade.Quantity * side
			} else {
				dateToPositions[date][trade.Instrument] = trade.Quantity * side
			}
		} else {
			dateToPositions[date] = map[string]int{trade.Instrument: trade.Quantity * side}
		}
	}
	return dateToPositions
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
	dateToPositionsCumulative := make(DateToPosition)
	dateToPositionsCumulative[positions[0].Date] = positions[0].Positions

	previous := positions[0].Positions
	for _, date := range positions[1:] {
		current := getCurrentSum(date.Positions, previous)
		dateToPositionsCumulative[date.Date] = current
		previous = current
		current = map[string]int{}
	}
	return dateToPositionsCumulative
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

func mapSingleDatePositionsToArray(positionsMap map[string]int) []Position {
	var positionsArray []Position
	for position, total := range positionsMap {
		positionsArray = append(positionsArray, Position{position, total})
	}
	sort.SliceStable(positionsArray, func(i, j int) bool {
		return positionsArray[i].Instrument < positionsArray[j].Instrument
	})

	return positionsArray
}
