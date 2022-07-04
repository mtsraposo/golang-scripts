package maths

import (
	"fmt"
	"math"
)

const MaxPoints = 1e2
const GamesToPlay = 1e6

func RunTennis(p float64) {
	firstWinsProb, deuceProb := SimulateGames(p, GamesToPlay)
	fmt.Printf("Probability of a Deuce: %f%% (vs. %f%% theoretical)\n", 100*deuceProb, 100*TheoreticalDeuceProb(p))
	fmt.Printf("Probability that the first player delivering will win: %f%% (vs. %f%% theoretical)\n",
		100*firstWinsProb, 100*TheoreticalFirstWinsProb(p))
}

func SimulateGames(p float64, n int) (float64, float64) {
	wins, deuces := 0.0, 0.0
	var winner int
	var deuce bool
	for i := 0; i < n; i++ {
		winner, deuce = SimulateGame(p)
		if winner == 1 {
			wins++
		}
		if deuce {
			deuces++
		}
	}
	return wins / float64(n), deuces / float64(n)
}

func SimulateGame(p float64) (int, bool) {
	score := [2]int{0, 0}
	deuce := false
	winner := 0
	for ; winner == 0; winner = whoWinsGame(score) {
		if Deuce(score) {
			deuce = true
		}
		if playerOneWinsPoint(p) {
			score = [2]int{score[0] + 1, score[1]}
		} else {
			score = [2]int{score[0], score[1] + 1}
		}
	}
	return winner, deuce
}

func Deuce(score [2]int) bool {
	return score[0] == 3 && score[1] == 3
}

func TheoreticalDeuceProb(p float64) float64 {
	return BinomialProb(6, 3, p)
}

func TheoreticalFirstWinsProb(p float64) float64 {
	prob := 0.0
	points := []int{3, 4, 5}
	for _, n := range points {
		prob += BinomialProb(n, 3, p) * p
	}

	pDeuce := TheoreticalDeuceProb(p)
	pWinAfterDeuce := math.Pow(p, 2) / (1 - BinomialProb(2, 1, p))
	prob += pDeuce * pWinAfterDeuce

	return prob
}
