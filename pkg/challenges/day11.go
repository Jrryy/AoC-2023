package challenges

import (
	"AoC-2023/pkg/utils"
)

func Day11() bool {
	input := utils.ImportInputMatrixChars(11)

	var coords, coords2 [][2]int

	for i := range input {
		for j, char := range input[i] {
			if char == '#' {
				coords = append(coords, [2]int{i, j})
				coords2 = append(coords2, [2]int{i, j})
			}
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if !utils.Contains(input[i], '#') {
			for _i, coord := range coords {
				if coord[0] > i {
					coords[_i] = [2]int{coord[0] + 1, coord[1]}
					coords2[_i] = [2]int{coords2[_i][0] + 999999, coords2[_i][1]}
				}
			}
		}
	}

	for j := len(input[0]) - 1; j >= 0; j-- {
		found := false
		for i := 0; i < len(input); i++ {
			if input[i][j] == '#' {
				found = true
				break
			}
		}
		if !found {
			for _i, coord := range coords {
				if coord[1] > j {
					coords[_i] = [2]int{coord[0], coord[1] + 1}
					coords2[_i] = [2]int{coords2[_i][0], coords2[_i][1] + 999999}
				}
			}
		}
	}

	total, total2 := 0, 0

	for i := range coords {
		for j := range coords[i:] {
			total += utils.Abs(coords[i][0]-coords[j+i][0]) + utils.Abs(coords[i][1]-coords[j+i][1])
			total2 += utils.Abs(coords2[i][0]-coords2[j+i][0]) + utils.Abs(coords2[i][1]-coords2[j+i][1])
		}
	}

	println(total)
	println(total2)

	return true
}
