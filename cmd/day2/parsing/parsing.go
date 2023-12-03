package parsing

import (
	"strings"
	"strconv"
	"adventOfCode23/cmd/day2/games" as Game
	
)

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
