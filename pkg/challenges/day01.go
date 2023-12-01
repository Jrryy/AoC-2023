package challenges

import (
	"AoC-2023/pkg/utils"
	"strings"
)

func Day01() {
	input := utils.ImportInputLines(1)

	total1, total2 := 0, 0

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range input {
		var firstDigit, firstIndex, lastDigit, lastIndex int

		for index, character := range line {
			if '0' <= character && character <= '9' {
				firstDigit = int(character) - int('0')
				firstIndex = index
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				lastDigit = int(line[i]) - int('0')
				lastIndex = i
				break
			}
		}
		total1 += firstDigit*10 + lastDigit

		firstSubstring := line[:firstIndex]
		for index, number := range numbers {
			found := strings.Index(firstSubstring, number)
			if found != -1 && found < firstIndex {
				firstDigit = index + 1
				firstIndex = found
			}
		}

		lastSubstring := line[lastIndex:]
		lastIndex = 0
		for index, number := range numbers {
			found := strings.LastIndex(lastSubstring, number)
			if found != -1 && found > lastIndex {
				lastDigit = index + 1
				lastIndex = found
			}
		}

		total2 += firstDigit*10 + lastDigit
	}

	println(total1)
	println(total2)
}
