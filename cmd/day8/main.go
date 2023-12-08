package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day8/parser"
	"fmt"
)

func main() {
	lines := common.ReadFileFromArgs()
	directions := parser.GetDirections(lines[0])
	theMap := parser.GetMap(lines[1:])
	
	found := false
	index := 0
	steps :=0
	location := "AAA"
	for {
		left := directions[index]
		
		location = GoNext(location, left, theMap)
		steps++

		if location == "ZZZ" {
			found = true
		}
		if index == len(directions) -1 {
			// Back to start
			index = 0
		} else {
			index = index + 1
		}

		if found {
			break
		}
	}

	fmt.Printf("\n it took %v steps to get to ZZZ\n", steps)


}

func GoNext(start string, left bool, theMap map[string][]string) string {
	options := theMap[start]
	if left {
		return options[0]
	}
	return options[1]
}
