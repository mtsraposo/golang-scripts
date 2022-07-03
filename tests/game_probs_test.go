package tests

import (
	. "github.com/mtsraposo/xp-strats/mathematics"
	. "github.com/mtsraposo/xp-strats/utils"
	"testing"
)

func TestGameProbs(t *testing.T) {
	tests := []struct {
		games     []*Node
		deuceProb float64
		winProb   float64
	}{
		{
			[]*Node{
				ArrayToLinkedList([][2]int{
					{0, 0}, {1, 0}, {1, 1}, {2, 1},
					{2, 2}, {2, 3}, {3, 3}, {4, 3},
					{4, 4}, {5, 4}, {5, 5}, {6, 5},
					{7, 5},
				}),
				ArrayToLinkedList([][2]int{
					{0, 0}, {1, 0}, {1, 1}, {2, 1},
					{2, 2}, {2, 3}, {3, 3}, {4, 4},
					{3, 5},
				}),
			},
			1.0, 0.5,
		},
	}
	for _, test := range tests {
		deuceProb, winProb := Probs(test.games)
		if deuceProb != test.deuceProb || winProb != test.winProb {
			t.Errorf("Failed to calculate probabilities. (deuceProb=%f, winProb=%f) does not equal (deuceProb=%f, winProb=%f)",
				deuceProb, winProb, test.deuceProb, test.winProb)
		}
	}
}
