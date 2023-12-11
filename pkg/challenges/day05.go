package challenges

import (
	"AoC-2023/pkg/utils"
	"math"
	"strconv"
	"strings"
)

func Day05() bool {
	input := strings.Split(strings.TrimRight(utils.ImportInput(5), "\n"), "\n\n")

	rawSeeds := strings.Split(input[0][7:], " ")

	seeds := make([]int, len(rawSeeds))
	seedRanges := make([][2]int, len(rawSeeds)/2)

	initial := true
	for i, rawSeed := range rawSeeds {
		seeds[i], _ = strconv.Atoi(rawSeed)
		if initial {
			seedRanges[i/2][0] = seeds[i]
		} else {
			seedRanges[i/2][1] = seeds[i]
		}
		initial = !initial
	}

	maps := make([][][3]int, 7)

	for i, rawInputMap := range input[1:] {
		inputMap := strings.Split(rawInputMap, "\n")
		maps[i] = make([][3]int, len(inputMap)-1)
		for j, mapNumbers := range inputMap[1:] {
			splitMapNumbers := strings.Split(mapNumbers, " ")
			maps[i][j][0], _ = strconv.Atoi(splitMapNumbers[0])
			maps[i][j][1], _ = strconv.Atoi(splitMapNumbers[1])
			maps[i][j][2], _ = strconv.Atoi(splitMapNumbers[2])
		}
	}

	minLocation1, minLocation2 := math.MaxInt, math.MaxInt

	for _, seed := range seeds {
		for _, mapper := range maps {
			for _, mapperNumbers := range mapper {
				if seed >= mapperNumbers[1] && seed <= mapperNumbers[1]+mapperNumbers[2]-1 {
					seed = mapperNumbers[0] + (seed - mapperNumbers[1])
					break
				}
			}
		}
		if seed < minLocation1 {
			minLocation1 = seed
		}
	}

	for _, seedRange := range seedRanges {
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]-1; seed++ {
			currentSeed := seed
			for _, mapper := range maps {
				for _, mapperNumbers := range mapper {
					if currentSeed >= mapperNumbers[1] && currentSeed <= mapperNumbers[1]+mapperNumbers[2]-1 {
						currentSeed = mapperNumbers[0] + (currentSeed - mapperNumbers[1])
						break
					}
				}
			}
			if currentSeed < minLocation2 {
				minLocation2 = currentSeed
			}
		}
	}

	println(minLocation1)
	println(minLocation2)

	return true
}
