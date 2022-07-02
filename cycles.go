package main

func permutationCycles(set []int) [][]int {
	positionMap := genPositionMap(set)
	visited := make(map[int]struct{})

	willVisitNext := func(pos int) (int, bool) {
		nextPosition, nextPosExists := positionMap[pos]
		_, nextPosVisited := visited[nextPosition]
		return nextPosition, nextPosExists && !nextPosVisited
	}

	var newCycle []int
	visit := func(posToVisit int) int {
		visited[posToVisit] = struct{}{}
		newCycle = append([]int{posToVisit}, newCycle...)
		return posToVisit
	}

	var cycles [][]int
	for i, _ := range set {
		startingPos := i + 1
		_, startingPosVisited := visited[startingPos]
		if !startingPosVisited {
			pos := visit(startingPos)
			for nextPosition, visitNext := willVisitNext(pos); visitNext; nextPosition, visitNext = willVisitNext(pos) {
				pos = visit(nextPosition)
			}
			cycles = append(cycles, rotate(newCycle, -1))
			newCycle = []int{}
		}
	}

	return cycles
}

func genPositionMap(set []int) map[int]int {
	positionMap := make(map[int]int)
	for i, num := range set {
		positionMap[i+1] = num
	}
	return positionMap
}
