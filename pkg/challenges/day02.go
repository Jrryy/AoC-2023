package challenges

import (
	"AoC-2023/pkg/utils"
	"strconv"
	"strings"
)

func Day02() bool {
	input := utils.ImportInputLines(2)

	total1, total2 := 0, 0

	for gameNumber, line := range input {
		pulls := strings.Split(line, ": ")[1]
		possible := true
		red, green, blue := 0, 0, 0
		for _, pull := range strings.Split(pulls, "; ") {
			for _, colourPull := range strings.Split(pull, ", ") {
				splitColourPull := strings.Split(colourPull, " ")
				switch splitColourPull[1] {
				case "red":
					number, _ := strconv.Atoi(splitColourPull[0])
					if number > red {
						red = number
					}
					if number > 12 {
						possible = false
					}
				case "green":
					number, _ := strconv.Atoi(splitColourPull[0])
					if number > green {
						green = number
					}
					if number > 13 {
						possible = false
					}
				case "blue":
					number, _ := strconv.Atoi(splitColourPull[0])
					if number > blue {
						blue = number
					}
					if number > 14 {
						possible = false
					}
				}
			}
		}
		if possible {
			total1 += gameNumber + 1
		}
		total2 += red * green * blue
	}

	println(total1)
	println(total2)
	return true
}
