package challenges

import (
	"AoC-2023/pkg/utils"
	"strings"
)

func Day04() bool {
	input := utils.ImportInputLines(4)

	cardCopies := make([]int, len(input))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	total1, total2 := 0, 0

	for cardNumber, card := range input {
		splitCard := strings.Split(card, ": ")
		numbers := splitCard[1]
		splitNumbers := strings.Split(numbers, " | ")
		myNumbers, winningNumbers := strings.Fields(splitNumbers[0]), strings.Fields(splitNumbers[1])
		points1, points2 := 0, 0
		for _, winningNumber := range winningNumbers {
			if utils.Contains(myNumbers, winningNumber) {
				points2++
				if points1 == 0 {
					points1 = 1
				} else {
					points1 *= 2
				}
			}
		}
		total1 += points1

		for i := cardNumber + 1; i < cardNumber+1+points2 && i < len(cardCopies); i++ {
			cardCopies[i] += cardCopies[cardNumber]
		}
		total2 += cardCopies[cardNumber]
	}

	println(total1)
	println(total2)

	return true
}
