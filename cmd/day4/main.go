package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day4/scratchcardreader"
	"fmt"
	"slices"
)

func main() {
	input := common.ReadFileFromArgs()
	totalScore := 0
	for gameNumber, line := range input {
		winners, gameValues := ReadScratchCard(line)
		score := CheckScoreOfGame(winners, gameValues)
		totalScore += score 
		fmt.Printf("\nScore for game %v was: %v\n", gameNumber+1, score)
	}
	fmt.Printf("\nTotal Score %v\n", totalScore)
}

func CheckScoreOfGame(game []int, winners []int) int {
	fmt.Printf("\nGame %v\nWinner %v\n", game, winners)
	score := 0
	for _, play := range game {
		if slices.Contains(winners, play) {
			score = AdjustScore(score)
		}
	} 

	// same loop reverse incase double values arent both scored as winners?

//	for _, winner := range winners {
//		if slices.Contains(game, winner) {
//			fmt.Printf("\nWinner: %v", winner)
//			score = AdjustScore(score)
//		}
//	}
	return score
}

func AdjustScore(score int) int {
	if score == 0 {
		score += 1
	} else {
		score *= 2
	}
	return score
}
