package tests

import (
	. "github.com/mtsraposo/golang-scripts/maths"
	"testing"
)

func TestGameProbs(t *testing.T) {
	tests := []struct {
		p float64
		n int
		winProb   float64
		deuceProb float64
	}{
		{0.7,1e6, 0.900108, 0.185171},
	}
	for _, test := range tests {
		winProb, deuceProb := SimulateGames(test.p, test.n)
		if deuceProb != test.deuceProb || winProb != test.winProb {
			t.Errorf("Failed to calculate probabilities. (winProb=%f, deuceProb=%f) does not equal (winProb=%f, deuceProb=%f)",
				winProb, deuceProb, test.winProb, test.deuceProb)
		} else {
			t.Logf("Probs(%f, %d) = (%f, %f)", test.p, test.n, winProb, deuceProb)
		}
	}
}
