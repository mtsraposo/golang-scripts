package scripts

func CalcMaxMinBreaches(prices []float64) [2]int {
	max, min := prices[0], prices[0]
	maxBreaches, minBreaches := 0, 0
	for _, price := range prices {
		if price > max {
			max = price
			maxBreaches++
		} else if price < min {
			min = price
			minBreaches++
		}
	}
	return [2]int{maxBreaches, minBreaches}
}
