package maths

import (
	"math"
	"math/rand"
)

func playerOneWinsPoint(p float64) bool {
	return rand.Float64() < p
}

func whoWinsGame(score [2]int) int {
	if playerOneWins := score[0] >= 4 && score[0]-score[1] >= 2; playerOneWins {
		return 1
	}
	if playerTwoWins := score[1] >= 4 && score[1]-score[0] >= 2; playerTwoWins {
		return 2
	}
	if score[0]+score[1] >= MaxPoints {
		return 3
	}
	return 0
}

func BinomialProb(n int, k int, p float64) float64 {
	combinations := float64(Factorial(n)) / float64(Factorial(n-k)*Factorial(k))
	return combinations * math.Pow(p, float64(k)) * math.Pow(1-p, float64(n-k))
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
