package main

import (
	"adventOfCode23/cmd/common"
	"strings"
	"strconv"
	"adventOfCode23/cmd/day2/games" as Game
	"adventOfCode23/cmd/day2/parsing"
)

func main() {
	var gamesStrings = common.ReadFileFromArgs()
	var bagOfCubes = initBagOfCubes()
	var games = ParseGames(gamesStrings)
	for _, game := range games {
		PrintGame(game)
	}
	println(bagOfCubes)
}

func PrintGame(game Game){
	println("\nGame:", game.Number)
	for setNum, set := range game.Sets {
		println("\n\tSet number:", setNum)
		for colour, count := range set.Colours {
			print(colour, ":", count, ",")
		}
	}
}

func initBagOfCubes() map[string]int {
	var bag = make(map[string]int)
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
	return bag
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

func ParseGames(lines []string) []Game {
	var games []Game;

	for index, line := range lines {
		var game = ParseGame(line)
		game.Number = index +1
		games = append(games, game)
	}
	return games
}

func ParseGame(line string) Game {
	var gameParts = strings.Split(line, ":")
	game := Game {
		Number: GetGameNumber(gameParts[0]), 
		Sets: make(map[int]Set), 
	}
	
	var setsSummary = strings.Split(gameParts[1], ";")
	for index, setLine := range setsSummary {
		game.Sets[index +1] = ParseSet(setLine)
	}
	
	return game
}

func GetGameNumber(line string) int {
	// todo
	return 1
}

func ParseSet(setString string) Set {
	set := Set { Colours: make(map[string]int)}
	var colourCounts = strings.Split(setString, ",")
	for _, draw := range colourCounts {
		draw = strings.TrimSpace(draw)
		var components = strings.Split(draw, " ")
		var count, err = strconv.Atoi(components[0])
		if (err != nil) {
			println("Parsing Error!", err)
		}
		set.Colours[components[1]] = count
	}
	return set
}

// Game Number:{set {Colour: count} }
type Game struct {
	Number int
	Sets map[int]Set
}

type Set struct {
	Colours map[string]int
}

