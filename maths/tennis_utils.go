package mathematics

import (
	"math"
	"math/rand"
)

func whoWinsPoint(p float64) int {
	outcome := rand.Float64()
	if outcome < p {
		return 1
	} else {
		return 2
	}
}

func whoWinsGame(score [2]int) int {
	if playerOneWins := score[0] >= 4 && score[0]-score[1] >= 2; playerOneWins {
		return 1
	}
	if playerTwoWins := score[1] >= 4 && score[1]-score[0] >= 2; playerTwoWins {
		return 2
	}
	return 0
}

func gameOver(score [2]int) bool {
	winner := whoWinsGame(score)
	tooManyPoints := score[0]+score[1] >= MaxPoints
	if winner != 0 || tooManyPoints {
		return true
	}
	return false
}

func firstWins(game *Node) bool {
	var winner int
	for node := game; node != nil; node = node.Next {
		winner = whoWinsGame(node.Score)
		if winner == 1 {
			return true
		}
		if winner == 2 {
			return false
		}
	}
	return false
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
