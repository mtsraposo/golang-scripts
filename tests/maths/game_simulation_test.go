package tests

import (
	. "github.com/mtsraposo/golang-scripts/maths"
	. "github.com/mtsraposo/golang-scripts/utils"
	"math/rand"
	"testing"
)

func TestSimulateGames(t *testing.T) {
	rand.Seed(42)
	tests := []struct {
		p     float64
		n     int
		games []*Node
	}{
		{0.7, 2,
			[]*Node{
				ArrayToLinkedList([][2]int{
					{0,0}, {1,0}, {2,0}, {3,0}, {4,0},
				}),
				ArrayToLinkedList([][2]int{
					{0,0}, {1,0}, {2,0}, {2,1}, {3,1}, {4,1},
				}),
			},
		},
	}
	for _, test := range tests {
		games := SimulateGames(test.p, test.n)
		if !GamesEqual(games, test.games) {
			t.Errorf("Simulated games do not equal for p = %f", test.p)
		}
		t.Logf("SimulateGames(%f, %d) = %v", test.p, test.n, LinkedListsToArrays(games))
	}
}
