package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day4/scratchcardreader"
	"fmt"
	"slices"
)

func main() {
	input := common.ReadFileFromArgs()
	scratchCards := make(map[int]int)
	for index := range input {
		gameNumber := index +1
		// handle iniial/original card
		count, exists := scratchCards[gameNumber]
		if !exists {
			// first time this card used
			scratchCards[gameNumber] = 1	
		} else {
			scratchCards[gameNumber] = count +1
		}
		playCard(gameNumber, input, scratchCards)

	}

	var totalCount int64 = 0
	for card, count := range scratchCards {
		fmt.Printf("Card %v had %v copies", card, count)
		totalCount += int64(count)
	}
	fmt.Printf("\nTotal Count %v\n", totalCount)
}

func CheckScoreOfGame(game []int, winners []int) int {
	fmt.Printf("\nGame %v\nWinner %v\n", game, winners)
	score := 0
	for _, play := range game {
		if slices.Contains(winners, play) {
			score += 1
		}
	} 

	return score
}

func playCard(gameNumber int, input []string, scratchCards map[int]int ) {
	index := gameNumber -1
	// Play a game as many times as it is in the scratchCards Map
	for i := 0; i < scratchCards[gameNumber]; i++ {
		fmt.Printf("\n Playing game %v", gameNumber)	
		winners, gameValues := ReadScratchCard(input[index])
		score := CheckScoreOfGame(winners, gameValues)
		fmt.Printf("\nScore for game %v was: %v\n", gameNumber, score)
		winCopysOfScratchCards(input, gameNumber, scratchCards, score)

	}
}

func winCopysOfScratchCards(allCards []string, currentGame int, scratchCards map[int]int, score int){
	for i:= 1; i <= score; i++ {
		nextCard := currentGame + i
		scratchCards[nextCard] = scratchCards[nextCard] +1
		fmt.Printf("\n now have %v copies of card %v\n", scratchCards[nextCard], nextCard)
	}
}
