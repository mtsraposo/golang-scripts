package utils

import (
	"errors"
	"time"
)

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

func IndexOf[T string | int](toFind T, arr []T) int {
	for i, elem := range arr {
		if elem == toFind {
			return i
		}
	}
	panic(errors.New("element not found"))
}

type FormattedTrade struct {
	TradeDate  string
	Instrument string
	SideIsBuy  bool
	Quantity   int
	Price      int
}

const LayoutISO = "2006-01-02"

func FormatDate(date time.Time) string {
	return date.Format(LayoutISO)
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
