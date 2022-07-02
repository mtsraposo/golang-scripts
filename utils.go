package main

func arraysEqual[V int | float32](A []V, B []V) bool {
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

func matricesEqual[V int | float32](A [][]V, B [][]V) bool{
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