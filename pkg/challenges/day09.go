package challenges

import (
	"AoC-2023/pkg/utils"
)

func allZeros(items []int) bool {
	for _, item := range items {
		if item != 0 {
			return false
		}
	}
	return true
}

func Day09() bool {
	input := utils.ImportInputLinesInts(9)

	total1, total2 := 0, 0

	for _, items := range input {
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
		var lastItems, firstItems []int

		for !allZeros(items) {
			lastItems = append(lastItems, items[0])
			firstItems = append(firstItems, items[len(items)-1])

			nextItems := make([]int, len(items)-1)

			for i := 0; i < len(items)-1; i++ {
				nextItems[i] = items[i] - items[i+1]
			}
			items = nextItems
		}

		continuation, previous := 0, 0

		for i := len(lastItems) - 1; i >= 0; i-- {
			continuation += lastItems[i]
			previous = firstItems[i] - previous
		}
		total1 += continuation
		total2 += previous
	}

	println(total1)
	println(total2)

	return true
}
