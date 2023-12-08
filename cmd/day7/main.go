package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day7/hand"
	"fmt"
	"sort"
)

func main(){
	lines := common.ReadFileFromArgs()
	hands := make([]Hand, len(lines))
	cardValues := MakeValueMap()

	for i, line := range lines {
		hands[i] = GetHand(line, cardValues)
		hands[i].Print()
	}
	PlayGame(hands)
}

func PlayGame(hands []Hand) {
	total := int64(0)
	hands = SortHands(hands)
	for rank, hand := range hands {
		win := int64((rank + 1) * hand.Bid)
		fmt.Printf("\n%v | %v * %v = %v",hand.String, rank +1, hand.Bid, win)
		total += win
	}
	
	fmt.Printf("\nEnd Winnings: %v\n", total)

}


func SortHands(hands []Hand) []Hand {
	var groups = make(map[int][]Hand )
	// Split into groups of same strength
	for _, hand := range hands {
		strength := hand.Strength
		groups[strength] = append(groups[strength], hand)
	}
	
	// sort child groups
	if len(groups) != len(hands) {
		for strength, childGroup := range groups {
			groups[strength] = Sort(childGroup, 0)
		}
	}
	return FinalSortLayer(groups, hands)
}

func Sort(hands []Hand, index int) []Hand {
	var groups = make(map[int][]Hand )
	// Split into groups of same strength
	for _, hand := range hands {
		strength := hand.HighCards[index]
		groups[strength] = append(groups[strength], hand)
	}
	
	// sort child groups
	if len(groups) != len(hands) {
		for strength, childGroup := range groups {
			groups[strength] = Sort(childGroup, index +1)
		}
	}
	return FinalSortLayer(groups, hands)	
}

func FinalSortLayer(groups map[int][]Hand, hands []Hand) []Hand {
	var strengths []int

	// Sort this layer
	for strength := range groups {
		strengths = append(strengths, strength)	
	}

	sort.Ints(strengths)
	
	groupsPreAssembly := make([][]Hand, len(groups))
	for i, strength := range strengths {
		groupsPreAssembly[i] = groups[strength]
	}

	hands = hands[:0]
	for _, group := range groupsPreAssembly {
		hands = append(hands, group...)
	}
	
//	fmt.Printf("\n Sorted Group to: %v", hands)
	return hands

}

func MakeValueMap() map[rune]int {
	values := make(map[rune]int)
	startScore := 2
	cards := []rune {'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

	for _, card := range cards {
		values[card] = startScore
		startScore++
	}
	return values
}
