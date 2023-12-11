package challenges

import (
	"AoC-2023/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

const cardOrder = "23456789TJQKA"
const cardOrder2 = "J23456789TQKA"

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards string
	bet   int
}

type handV2 struct {
	hand
}

func (h hand) evaluateHand() int {
	cardRunes := []rune(h.cards)
	mappedCards := make(map[rune]int)
	for _, char := range cardRunes {
		mappedCards[char]++
	}
	if len(mappedCards) == 1 {
		return fiveOfAKind
	}
	if len(mappedCards) == 2 {
		if mappedCards[cardRunes[0]] == 1 || mappedCards[cardRunes[0]] == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	}
	if len(mappedCards) == 3 {
		if mappedCards[cardRunes[0]] == 3 || mappedCards[cardRunes[1]] == 3 || mappedCards[cardRunes[2]] == 3 {
			return threeOfAKind
		} else {
			return twoPair
		}
	}
	if len(mappedCards) == 4 {
		return onePair
	}
	return highCard
}

func (h handV2) evaluateHandV2() int {
	allCardRunes := []rune(h.cards)
	var cardRunes []rune
	mappedCards := make(map[rune]int)
	for _, char := range allCardRunes {
		mappedCards[char]++
		if char != 'J' {
			cardRunes = append(cardRunes, char)
		}
	}
	jokers := mappedCards['J']
	delete(mappedCards, 'J')
	// fmt.Println(mappedCards)
	if len(mappedCards) == 1 || jokers == 5 {
		return fiveOfAKind
	}
	if len(mappedCards) == 2 {
		if mappedCards[cardRunes[0]]+jokers == 4 || mappedCards[cardRunes[1]]+jokers == 4 {
			return fourOfAKind
		} else {
			return fullHouse
		}
	}
	if len(mappedCards) == 3 {
		if mappedCards[cardRunes[0]]+jokers == 3 || mappedCards[cardRunes[1]]+jokers == 3 || mappedCards[cardRunes[2]]+jokers == 3 {
			return threeOfAKind
		} else {
			return twoPair
		}
	}
	if len(mappedCards) == 4 {
		return onePair
	}
	return highCard
}

type handsList []hand

func (hl handsList) Len() int {
	return len(hl)
}

func (hl handsList) Less(i, j int) bool {
	iPlay, jPlay := hl[i].evaluateHand(), hl[j].evaluateHand()
	if iPlay != jPlay {
		return iPlay < jPlay
	}
	iCards, jCards := strings.Split(hl[i].cards, ""), strings.Split(hl[j].cards, "")
	for i := range iCards {
		if strings.Index(cardOrder, iCards[i]) != strings.Index(cardOrder, jCards[i]) {
			return strings.Index(cardOrder, iCards[i]) < strings.Index(cardOrder, jCards[i])
		}
	}
	return true
}

func (hl handsList) Swap(i, j int) {
	hl[i], hl[j] = hl[j], hl[i]
}

type handsListV2 []handV2

func (hl handsListV2) Len() int {
	return len(hl)
}

func (hl handsListV2) Less(i, j int) bool {
	iPlay, jPlay := hl[i].evaluateHandV2(), hl[j].evaluateHandV2()
	if iPlay != jPlay {
		return iPlay < jPlay
	}
	iCards, jCards := strings.Split(hl[i].cards, ""), strings.Split(hl[j].cards, "")
	for i := range iCards {
		if strings.Index(cardOrder2, iCards[i]) != strings.Index(cardOrder2, jCards[i]) {
			return strings.Index(cardOrder2, iCards[i]) < strings.Index(cardOrder2, jCards[i])
		}
	}
	return true
}

func (hl handsListV2) Swap(i, j int) {
	hl[i], hl[j] = hl[j], hl[i]
}

func Day07() bool {
	input := utils.ImportInputLines(7)

	hands := make(handsList, len(input))
	handsV2 := make(handsListV2, len(input))

	for i, line := range input {
		splitLine := strings.Split(line, " ")
		bet, _ := strconv.Atoi(splitLine[1])
		hands[i].cards = splitLine[0]
		hands[i].bet = bet
		handsV2[i].cards = splitLine[0]
		handsV2[i].bet = bet
	}

	sort.Sort(hands)
	sort.Sort(handsV2)

	total1, total2 := 0, 0

	for i := range hands {
		total1 += (i + 1) * hands[i].bet
		total2 += (i + 1) * handsV2[i].bet
	}

	println(total1)
	println(total2)

	return true
}
