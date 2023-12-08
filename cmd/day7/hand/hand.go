package hand

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strings"
)

type Hand struct {
	String string
	Totals map[rune]int
	Bid int
	Strength int
	HighCards []int
}

func GetHand(line string, highCardMap map[rune]int) Hand {
	parts := strings.Split(line, " ")
	bid := common.StringToInt(parts[1])
	hand := strings.TrimSpace(parts[0])
	totals := make(map[rune]int)
	for _, char := range hand {
		totals[char] = totals[char] +1
	}
	totals = Jokerify(totals)
	strength := GetStrength(totals)
	highCardScores := CalculateHighCards(hand, highCardMap)
	return Hand {
		String : hand,
		Totals: totals,
		Bid : bid,
		Strength: strength,
		HighCards: highCardScores,
	}
}

func (h Hand) Print() {
	fmt.Printf("\nHand was: %v, which has the totals %v, a bid of %v, and a strength of %v", h.String, h.Totals, h.Bid, h.Strength)
}

func GetStrength(totals map[rune]int) int {
	numberOfCards := len(totals)
	switch {
	case numberOfCards == 1 :
		println("Five of a kind")
		return 7
	case numberOfCards == 2 :
		// 4 of a kind or full house
		for _, count := range totals {
			if count == 4 || count == 1 {
				// 4 of a kind
				return 6
			} else {
				// full house
				return 5
			}
		}
	case numberOfCards == 3 : 
		// two pairs or three of a kind	
		for _, count := range totals {
			if count == 2 {
				// Two Pairs
				return 3
			} 
			if count == 3{
				// Three Of a Kind
				return 4
			}
		}
	case numberOfCards == 4 :
		// A pair
		return 2
	default : 
		return 1
	}
	fmt.Println("Should never hit here")
	return 0
}

func CalculateHighCards(input string, highCardMap map[rune]int) (scores []int) {
	scores = make([]int, len(input))
	for i, char := range input {
//		fmt.Printf("\nChar %v = Score %v ", string(rune(char)), highCardMap[rune(char)])
		scores[i] = highCardMap[rune(char)]
	}
	//fmt.Println(" High Card Scores : ", scores)
	return scores
}

func Jokerify(cards map[rune]int) map[rune]int {
	jokers, any := cards['J']
	if !any || jokers == 5 {
		fmt.Printf("\nNo Jokes for")
		for r := range cards {
		fmt.Print(string(rune(r)))
		}
		return cards
	}
	newCards := make(map[rune]int)
	numberOfCards := len(cards)

	switch {
	case numberOfCards == 2 :
		// Any change will make a 5 of a Kind
		for card, _ := range cards {
			if card != 'J' {
				newCards[card] = 5
			}
		}
		return newCards

	case numberOfCards == 3 && jokers == 2 :
		for card, count := range cards {
			if card != 'J' && count == 2 {
				newCards[card] = 4
			} else if card != 'J' {
				newCards[card] = count
			}
		}
		return newCards
	case numberOfCards == 3 :
		// 2 X 2 -> Full house
		// 3 of a kind -> 4 of a kind
		alreadyJockered := false
		for card, count := range cards {
			if count == 3 {
				newCards[card] = 4
			} else if count == 2 && !alreadyJockered {
				newCards[card] = 3
				alreadyJockered = true
			} else if card != 'J' {
				newCards[card] = count
			}
		}
		return newCards
	case numberOfCards == 4 && jokers == 2 :
		alreadyJockered := false
		for card, count := range cards {
			if !alreadyJockered && card != 'J' {
				newCards[card] = 3
				alreadyJockered = true
			} else if card != 'J' {
				newCards[card] = count
			}
		}
		return newCards


	case numberOfCards == 4:
		// Pair becomes 3 of a kind
		for card, count := range cards {
			if count == 2 {
				newCards[card] = 3
			} else if card != 'J' {
				newCards[card] = count
			}
		}
		return newCards
	case numberOfCards == 5:
		alreadyJockered := false
		for card, count := range cards {
			if !alreadyJockered && card != 'J' {
				newCards[card] = 2
				alreadyJockered = true
			} else if card != 'J' {
				newCards[card] = count
			}
		}
		return newCards
	}
	fmt.Println("!-- Should Never hit here--!")
	return cards
}

