package tests

import (
	. "github.com/mtsraposo/xp-strats/mathematics"
	"testing"
)

func TestBinomialProb(t *testing.T) {
	tests := []struct {
		n       int
		k       int
		p       float64
		binProb float64
	}{
		{1, 1, 0.5, 0.5},
		{2, 1, 0.5, 0.5},
		{3, 2, 0.5, 0.375},
		{6, 3, 0.5, 0.3125},
	}
	for _, test := range tests {
		binProb := BinomialProb(test.n, test.k, test.p)
		if binProb != test.binProb {
			t.Errorf("Binomial probabilities do not match for (n=%d, k=%d, p=%f). %f does not equal %f",
				test.n, test.k, test.p, binProb, test.binProb)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		fact int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}
	for _, test := range tests {
		fact := Factorial(test.n)
		if fact != test.fact {
			t.Errorf("Factorials do not match for number %d. %d does not equal %d",
				test.n, fact, test.fact)
		}
	}
}
