package main

func arraysEqual[V int | float64](A []V, B []V) bool {
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

func matricesEqual[V int | float64](A [][]V, B [][]V) bool{
	if len(A) != len(B) {
		return false
	}
	for i, row := range B {
		if !arraysEqual(row, A[i]) {
			return false
		}
	}
	return true
}

func min(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}