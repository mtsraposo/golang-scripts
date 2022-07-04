package maths

import (
	"fmt"
	"math"
)

type Node struct {
	Score [2]int
	Next  *Node
}

const MaxPoints = 1e2
const GamesToPlay = 1e6

func RunTennis(p float64) {
	games := SimulateGames(p, GamesToPlay)
	deuceProb, firstWinsProb := Probs(games)
	fmt.Printf("Probability of a Deuce: %f%% (vs. %f%% theoretical)\n", 100*deuceProb, 100*TheoreticalDeuceProb(p))
	fmt.Printf("Probability that the first player delivering will win: %f%% (vs. %f%% theoretical)\n",
		100*firstWinsProb, 100*TheoreticalFirstWinsProb(p))
}

func SimulateGames(p float64, n int) []*Node {
	var games []*Node
	for i := 0; i < n; i++ {
		game := Node{[2]int{0, 0}, nil}
		SimulateGame(&game, p)
		games = append(games, &game)
	}
	return games
}

func SimulateGame(root *Node, p float64) {
	if gameOver(root.Score) {
		return
	}

	if whoWinsPoint(p) == 1 {
		won := [2]int{root.Score[0] + 1, root.Score[1]}
		root.Next = &Node{won, nil}
	} else {
		lost := [2]int{root.Score[0], root.Score[1] + 1}
		root.Next = &Node{lost, nil}
	}
	SimulateGame(root.Next, p)
}

func Probs(games []*Node) (float64, float64) {
	deuces, wins := 0.0, 0.0
	for _, game := range games {
		if Deuce(game) {
			deuces++
		}
		if firstWins(game) {
			wins++
		}
	}
	gamesPlayed := float64(len(games))
	return deuces / gamesPlayed, wins / gamesPlayed
}

func Deuce(game *Node) bool {
	for node := game; node != nil; node = node.Next {
		if node.Score[0] == 3 && node.Score[1] == 3 {
			return true
		}
	}
	return false
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