package games

type Game struct {
	Number int
	Sets map[int]Set
}

type Set struct {
	Colours map[string]int
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
