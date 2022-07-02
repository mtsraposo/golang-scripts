package main

func rotate(L []int, k int) []int {
	n := len(L)
	pivotPosition := _calcPivotPosition(k, n)
	if pivotPosition != 0 {
		return append(L[pivotPosition:], L[:pivotPosition]...)
	} else {
		return L
	}
}

func _calcPivotPosition(k int, n int) int {
	pivotPosition := k % n
	if pivotPosition >= 0 {
		return pivotPosition
	} else {
		return n + pivotPosition
	}
}
