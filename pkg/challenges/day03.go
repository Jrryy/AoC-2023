package challenges

import (
	"AoC-2023/pkg/utils"
	"strconv"
)

func checkSurroundings(input [][]rune, i, j int) (bool, bool, int, int) {
	if i > 0 {
		if j > 0 {
			if !(input[i-1][j-1] >= '0' && input[i-1][j-1] <= '9') && input[i-1][j-1] != '.' {
				isGear := input[i-1][j-1] == '*'
				return true, isGear, i - 1, j - 1
			}
		}
		if j < len(input[0])-1 {
			if !(input[i-1][j+1] >= '0' && input[i-1][j+1] <= '9') && input[i-1][j+1] != '.' {
				isGear := input[i-1][j+1] == '*'
				return true, isGear, i - 1, j + 1
			}
		}
		if !(input[i-1][j] >= '0' && input[i-1][j] <= '9') && input[i-1][j] != '.' {
			isGear := input[i-1][j] == '*'
			return true, isGear, i - 1, j
		}
	}
	if i < len(input)-1 {
		if j > 0 {
			if !(input[i+1][j-1] >= '0' && input[i+1][j-1] <= '9') && input[i+1][j-1] != '.' {
				isGear := input[i+1][j-1] == '*'
				return true, isGear, i + 1, j - 1
			}
		}
		if j < len(input[0])-1 {
			if !(input[i+1][j+1] >= '0' && input[i+1][j+1] <= '9') && input[i+1][j+1] != '.' {
				isGear := input[i+1][j+1] == '*'
				return true, isGear, i + 1, j + 1
			}
		}
		if !(input[i+1][j] >= '0' && input[i+1][j] <= '9') && input[i+1][j] != '.' {
			isGear := input[i+1][j] == '*'
			return true, isGear, i + 1, j
		}
	}
	if j > 0 {
		if !(input[i][j-1] >= '0' && input[i][j-1] <= '9') && input[i][j-1] != '.' {
			isGear := input[i][j-1] == '*'
			return true, isGear, i, j - 1
		}
	}
	if j < len(input[0])-1 {
		if !(input[i][j+1] >= '0' && input[i][j+1] <= '9') && input[i][j+1] != '.' {
			isGear := input[i][j+1] == '*'
			return true, isGear, i, j + 1
		}
	}
	return false, false, 0, 0
}

func Day03() bool {
	input := utils.ImportInputMatrixChars(3)

	total1, total2 := 0, 0

	currentDigits := ""
	isPart := false
	isGear := false
	var iGear, jGear int

	gears := make(map[[2]int][]int)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] >= '0' && input[i][j] <= '9' {
				currentDigits += string(input[i][j])
				currentIsPart, currentIsGear, iPart, jPart := checkSurroundings(input, i, j)
				if !isPart && currentIsPart {
					isPart = true
					if currentIsGear {
						isGear = true
						iGear, jGear = iPart, jPart
					}
				}
			} else {
				if isPart {
					currentNumber, _ := strconv.Atoi(currentDigits)
					total1 += currentNumber
					if isGear {
						gears[[2]int{iGear, jGear}] = append(gears[[2]int{iGear, jGear}], currentNumber)
					}
				}
				isPart = false
				isGear = false
				currentDigits = ""
			}
		}
		if isPart {
			currentNumber, _ := strconv.Atoi(currentDigits)
			total1 += currentNumber
			if isGear {
				gears[[2]int{iGear, jGear}] = append(gears[[2]int{iGear, jGear}], currentNumber)
			}
		}
		isPart = false
		isGear = false
		currentDigits = ""
	}

	for _, gearsArray := range gears {
		if len(gearsArray) == 2 {
			total2 += gearsArray[0] * gearsArray[1]
		}
	}

	println(total1)
	println(total2)
	return true
}
