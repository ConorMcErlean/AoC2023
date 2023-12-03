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
	var possibleGames, powersOfBags = CheckGames(bagOfCubes, games)
	fmt.Printf("\nPossible Games: %v , powers of min bags %v", possibleGames, powersOfBags)
	var totalOfPossible, sumOfPowers = 0, 0
	for _, gameNumber := range possibleGames {
		totalOfPossible += gameNumber
	}
	for _, power := range powersOfBags {
		sumOfPowers += power
	}
	println("\nThe total of the possible games was:", totalOfPossible )
	println("\nThe totalPowes of the possible games min bags was:", sumOfPowers )
}

func initBagOfCubes() map[string]int {
	var bag = make(map[string]int)
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
	return bag
}

func CheckGames(bag map[string]int, games []Game ) (possibleGames []int, powerOfMinBags []int) {
	for _, game := range games {
		var impossibleGames []int
		if WasGamePossible(game, bag){
			fmt.Println("Game", game.Number, " was possible")
			possibleGames = append(possibleGames, game.Number)
		} else {
			fmt.Println("Game", game.Number, " was not possible")
			impossibleGames = append(impossibleGames, game.Number)
		}

		minBag := CheckMinimumPossible(game)
		powerOfMinBags = append(powerOfMinBags, GetBagPower(minBag))
		fmt.Println("Minimum required would have been:", minBag)
	}
	// fmt.Printf("\nPossible Games: %v, \nImpossible games: %v", possibleGames, impossibleGames)
	return possibleGames, powerOfMinBags
}

func GetBagPower(bag map[string]int) int {
	// Initialise to 1 as 1*X = X
	var power = 1
	for _, count := range bag {
		power *= count
	}
	return power
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

func CheckMinimumPossible(game Game) (bag map[string]int) {
	bag = make(map[string]int)
	for _, set := range game.Sets {
		// Basically for a set take the highest number of a colour, and set the bag as that
		for colour, count := range set.Colours {
			if count > bag[colour] {
				bag[colour] = count
			}
		}
	}
	return bag
}

