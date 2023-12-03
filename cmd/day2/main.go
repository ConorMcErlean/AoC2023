package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day2/games"
	. "adventOfCode23/cmd/day2/parsing"
	"fmt"
)

func main() {
	var gamesStrings = common.ReadFileFromArgs()
	var bagOfCubes = initBagOfCubes()
	var games = ParseGames(gamesStrings)
	var possibleGames, impossibleGames = CheckGames(bagOfCubes, games)
	fmt.Printf("\nPossible Games: %v, \nImpossible games: %v", possibleGames, impossibleGames)
	var totalOfPossible = 0
	for _, gameNumber := range possibleGames {
		totalOfPossible += gameNumber
	}
	println("\nThe total of the possible games was:", totalOfPossible )
}

func initBagOfCubes() map[string]int {
	var bag = make(map[string]int)
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
	return bag
}

func CheckGames(bag map[string]int, games []Game ) (possibleGames []int, impossibleGames []int) {
	for _, game := range games {
		if WasGamePossible(game, bag){
			fmt.Println("Game", game.Number, " was possible")
			possibleGames = append(possibleGames, game.Number)
		} else {
			fmt.Println("Game", game.Number, " was not possible")
			impossibleGames = append(impossibleGames, game.Number)
		}
	}
	return possibleGames, impossibleGames
}

func WasGamePossible(game Game, bag map[string]int) bool {
	for _, set := range game.Sets {
		for colour, count := range set.Colours {
			if !checkIfPossible(bag, colour, count) {
				return false 
			}
		}
	}
	return true
}

func checkIfPossible(bag map[string]int, colour string, number int) bool {
	_, exists := bag[colour]
	if !exists {
		return false
	}
	var totalPossible = bag[colour]
	if number > totalPossible {
		return false
	}
	return true
}

