package scripts

import (
	"math"
)

const LEFT = 0
const RIGHT = 1

func union(intervals [][]float64) [][]float64 {
	return mergeSort(intervals, 0, len(intervals))
}

func mergeSort(intervals [][]float64, left int, right int) [][]float64 {
	midpoint := (left + right) / 2
	if left == midpoint {
		return [][]float64{intervals[left]}
	}
	leftUnion := mergeSort(intervals, left, midpoint)
	rightUnion := mergeSort(intervals, midpoint, right)
	//fmt.Printf("%v + %v", leftUnion, rightUnion)
	return merge(leftUnion, rightUnion)
}

func merge(a [][]float64, b [][]float64) [][]float64 {
	aSize, bSize := len(a), len(b)
	n := aSize + bSize

	left, right, mergeSize := 1, 0, 1
	side := RIGHT
	current := [][]float64{a[0]}
	merged := current
	for i := 1; i < n; i++ {
		if right == bSize || left < aSize && side == LEFT {
			current, side = mergePair(merged[mergeSize-1], a[left])
			left++
		} else {
			current, side = mergePair(merged[mergeSize-1], b[right])
			right++
		}
		merged = append(merged[:mergeSize-1], current...)
		mergeSize = len(merged)
	}
	//fmt.Printf(" = %v\n", merged)
	return merged
}

func mergePair(aUnit []float64, bUnit []float64) ([][]float64, int) {
	if aUnit[1] < bUnit[0] {
		return [][]float64{aUnit, bUnit}, LEFT
	}
	if aUnit[0] > bUnit[1] {
		return [][]float64{bUnit, aUnit}, RIGHT
	}
	start := math.Min(aUnit[0], bUnit[0])
	end := math.Max(aUnit[1], bUnit[1])
	if aUnit[1] > bUnit[1] {
		return [][]float64{{start, end}}, RIGHT
	} else {
		return [][]float64{{start, end}}, LEFT
	}
}
