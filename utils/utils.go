package utils

import (
	. "github.com/mtsraposo/xp-strats/mathematics"
)

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
