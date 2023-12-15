package challenges

import (
	"AoC-2023/pkg/utils"
)

func Day10() bool {
	matrix := utils.ImportInputMatrixChars(10)
	var pathMatrix [140][140]bool

	// I'm just gonna customize all of this with my input
	lastI, lastJ := 21, 29
	pathMatrix[lastI][lastJ] = true
	i, j := lastI-1, lastJ
	steps := 0

	for matrix[i][j] != 'S' {
		pathMatrix[i][j] = true
		switch matrix[i][j] {
		case '|':
			if lastI == i-1 {
				i += 1
				lastI += 1
			} else {
				i -= 1
				lastI -= 1
			}
		case 'J':
			if lastI == i-1 {
				j -= 1
				lastI += 1
			} else {
				i -= 1
				lastJ += 1
			}
		case 'L':
			if lastI == i-1 {
				j += 1
				lastI += 1
			} else {
				i -= 1
				lastJ -= 1
			}
		case 'F':
			if lastI == i+1 {
				j += 1
				lastI -= 1
			} else {
				i += 1
				lastJ -= 1
			}
		case '7':
			if lastI == i+1 {
				j -= 1
				lastI -= 1
			} else {
				i += 1
				lastJ += 1
			}
		case '-':
			if lastJ == j-1 {
				j += 1
				lastJ += 1
			} else {
				j -= 1
				lastJ -= 1
			}
		}
		steps++
	}

	println(steps/2 + 1)

	matrix[i][j] = 'L'

	inside := 0

	for i, row := range pathMatrix {
		for j, loop := range row {
			crossedWalls := 0
			if !loop {
				fromBelow, fromAbove := false, false
				for _j := j; _j < len(row); _j++ {
					if pathMatrix[i][_j] {
						if matrix[i][_j] == '|' {
							crossedWalls++
						} else if matrix[i][_j] == 'L' {
							fromAbove = true
						} else if matrix[i][_j] == 'F' {
							fromBelow = true
						} else if matrix[i][_j] == 'J' {
							if fromBelow {
								crossedWalls++
							}
							fromAbove = false
							fromBelow = false
						} else if matrix[i][_j] == '7' {
							if fromAbove {
								crossedWalls++
							}
							fromAbove = false
							fromBelow = false
						}
					}
				}
				if crossedWalls%2 == 1 {
					inside++
				}
			}
		}
	}

	println(inside)

	return true
}
