package challenges

import (
	"AoC-2023/pkg/utils"
)

func endReached(nodes, endingNodes []string) bool {
	for _, node := range nodes {
		if !utils.Contains(endingNodes, node) {
			return false
		}
	}
	return true
}

func Day08() bool {
	input := utils.ImportInputLines(8)

	instructions := []rune(input[0])
	instructionsSize := len(instructions)

	tree := make(map[string][2]string)
	var startingNodes, endingNodes []string

	for _, line := range input[2:] {
		tree[line[0:3]] = [2]string{line[7:10], line[12:15]}
		if line[2:3] == "A" {
			startingNodes = append(startingNodes, line[0:3])
		}
		if line[2:3] == "Z" {
			endingNodes = append(endingNodes, line[0:3])
		}
	}

	currentNode := "AAA"
	step := 0

	for currentNode != "ZZZ" {
		instruction := instructions[step%instructionsSize]
		if instruction == 'L' {
			currentNode = tree[currentNode][0]
		} else {
			currentNode = tree[currentNode][1]
		}
		step++
	}

	println(step)

	stepsEachPath := make([]int, len(startingNodes))
	currentNodes := startingNodes
	for i := 0; ; i++ {
		for j, node := range currentNodes {
			if utils.Contains(endingNodes, node) {
				stepsEachPath[j] = i
			}
			instruction := instructions[i%instructionsSize]
			if instruction == 'L' {
				currentNodes[j] = tree[node][0]
			} else {
				currentNodes[j] = tree[node][1]
			}
		}
		if !utils.Contains(stepsEachPath, 0) {
			break
		}
	}

	a := stepsEachPath[0]
	for _, b := range stepsEachPath[1:] {
		a = utils.LCM(a, b)
	}

	println(a)

	return true
}
