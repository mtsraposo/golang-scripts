package maths

import "fmt"

type BallotProb struct {
	fromBlue float64
	fromRed  float64
	total    float64
}

func RunBallot(red int, blue int) {
	fmt.Printf("P(last ball removed is blue) = %f%%\n", 100*ProbLastRemovedIsBlue(red, blue))
}

func ProbLastRemovedIsBlue(red int, blue int) float64 {
	dp := make([][]BallotProb, red)

	var i, j int
	for i = 0; i < red; i++ {
		dp[i] = make([]BallotProb, blue)
		dp[i][0] = BallotProb{0, 0, 0}
	}
	for j = 0; j < blue; j++ {
		dp[0][j] = BallotProb{1, 1, 1}
	}

	var toBlue, toRed, total, fromBlue, fromRed, reds, blues float64
	for i = 1; i < red; i++ {
		for j = 1; j < blue; j++ {
			reds, blues = float64(i) + 1, float64(j) + 1
			toBlue = blues / (reds + blues) * dp[i][j-1].fromBlue
			toRed = reds / (reds + blues) * dp[i-1][j].fromRed
			total = toBlue + toRed
			fromBlue = (reds*total + blues*dp[i][j-1].fromBlue) / (reds + blues)
			fromRed = (reds*dp[i-1][j].fromRed + blues*total) / (reds + blues)
			dp[i][j] = BallotProb{fromBlue, fromRed, total}
		}
	}

	return dp[red-1][blue-1].total
}
