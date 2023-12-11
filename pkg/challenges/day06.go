package challenges

import (
	"AoC-2023/pkg/utils"
	"strconv"
	"strings"
)

func Day06() bool {
	input := utils.ImportInputLines(6)

	rawTimeRecords := strings.Fields(input[0][5:])
	rawDistances := strings.Fields(input[1][10:])

	var timeRecords, distances []int
	singleTimeRecord, _ := strconv.Atoi(strings.ReplaceAll(input[0][5:], " ", ""))
	singleDistance, _ := strconv.Atoi(strings.ReplaceAll(input[1][10:], " ", ""))

	for _, timeRecord := range rawTimeRecords {
		newTimeRecord, _ := strconv.Atoi(timeRecord)
		timeRecords = append(timeRecords, newTimeRecord)
	}

	for _, distance := range rawDistances {
		newTimeRecord, _ := strconv.Atoi(distance)
		distances = append(distances, newTimeRecord)
	}

	total := 1

	for i := 0; i < len(timeRecords); i++ {
		options := 0
		for j := 1; j < timeRecords[i]; j++ {
			if j*(timeRecords[i]-j) > distances[i] {
				options++
			}
		}
		total *= options
	}

	maxTime, minTime := singleTimeRecord, 0
	current := singleTimeRecord / 2
	for minTime != maxTime-1 {
		if current*(singleTimeRecord-current) > singleDistance {
			maxTime = current
		} else {
			minTime = current
		}
		current = (maxTime + minTime) / 2
	}

	println(total)
	println(singleTimeRecord - current*2 - 1)

	return true
}
