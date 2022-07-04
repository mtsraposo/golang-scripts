package utils

import (
	. "github.com/mtsraposo/golang-scripts/maths"
	. "github.com/mtsraposo/golang-scripts/scripts"
	"time"
)

func PositionsEqual(p1 []Position, p2 []Position) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i, pos := range p2 {
		if pos.Instrument != p2[i].Instrument || pos.Total != p2[i].Total {
			return false
		}
	}
	return true
}

func ArraysEqual[V int | float64](A []V, B []V) bool {
	if len(A) != len(B) {
		return false
	}

	for i, v := range A {
		if v != B[i] {
			return false
		}
	}

	return true
}

func MatricesEqual[V int | float64](A [][]V, B [][]V) bool {
	if len(A) != len(B) {
		return false
	}
	for i, row := range B {
		if !ArraysEqual(row, A[i]) {
			return false
		}
	}
	return true
}

func Min(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func GamesEqual(games1 []*Node, games2 []*Node) bool {
	if len(games1) != len(games2) {
		return false
	}
	for i, g1 := range games1 {
		node1 := g1
		node2 := games2[i]
		for node1 != nil && node2 != nil {
			scoresDiffer := node1.Score != node2.Score
			if scoresDiffer {
				return false
			}
			node1 = node1.Next
			node2 = node2.Next
		}
		differentLens := (node1 == nil && node2 != nil) || (node2 == nil && node1 != nil)
		if differentLens {
			return false
		}
	}
	return true
}

func ArrayToLinkedList(scores [][2]int) *Node {
	if len(scores) == 0 {
		return &Node{}
	}
	root := &Node{Score: scores[0]}
	node := root
	for _, score := range scores[1:] {
		node.Next = &Node{Score: score}
		node = node.Next
	}
	return root
}

func LinkedListsToArrays(games []*Node) [][][2]int {
	var gamesArray [][][2]int
	var gameArray [][2]int
	for _, game := range games {
		node := game
		gameArray = [][2]int{}
		for node.Next != nil {
			gameArray = append(gameArray, node.Score)
			node = node.Next
		}
		gamesArray = append(gamesArray, gameArray)
	}
	return gamesArray
}

type FormattedTrade struct {
	TradeDate  string
	Instrument string
	SideIsBuy  bool
	Quantity   int
	Price      int
}

func FormatTrades(trades []Trade) []FormattedTrade {
	var fmtTrades []FormattedTrade
	var fmtTrade FormattedTrade
	for _, trade := range trades {
		fmtTrade = FormattedTrade{
			FormatDate(trade.TradeDate),
			trade.Instrument,
			trade.SideIsBuy,
			trade.Quantity,
			trade.Price,
		}
		fmtTrades = append(fmtTrades, fmtTrade)
	}
	return fmtTrades
}

const layoutISO = "2006-01-02"

func FormatDate(date time.Time) string {
	return date.Format(layoutISO)
}
