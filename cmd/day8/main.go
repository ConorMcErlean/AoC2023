package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day8/parser"
	"fmt"
	"strings"
)

func main() {
	lines := common.ReadFileFromArgs()
	directions := parser.GetDirections(lines[0])
	theMap := parser.GetMap(lines[1:])
	GoLikeAGhostEfficiently(directions, theMap)
	
}

func GoToEnd(directions []bool, theMap map[string][]string){
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

func GoLikeAGhost(directions []bool, theMap map[string][]string){
	found := false
	index := 0
	steps := int64(0)
	locations := GetGhostStarts(theMap)
	for {
		left := directions[index]
		for index, location := range locations {
			locations[index] = GoNext(location, left, theMap)
		}
		steps++
		fmt.Printf("\nStep %v", steps)

		if CheckGhostFinish(locations) {
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

func GoLikeAGhostEfficiently(directions []bool, theMap map[string][]string){
	// See how long each ghostStart takes to get
	locations := GetGhostStarts(theMap)
	cycleLengths := make([]int, len(locations))
	for index, location := range locations {
		cycleLengths[index] = FindCycleLength(location, directions, theMap)
	}
	otherCycles := cycleLengths[2:]
	fmt.Printf("Cycle Lengths of %v", cycleLengths)
	lowestCommon := FindLowestCommonMultiple(cycleLengths[0], cycleLengths[1], otherCycles...)
	fmt.Printf("\nLowest Common %v\n", lowestCommon)

}

func FindLowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = FindLowestCommonMultiple(result, integers[i])
	}

	return result

}

func GreatestCommonDivisor(a, b int) int {
	// Keep Modulusing till no remainder
	for b != 0 {
		x := b
		b = a % b
		a = x 
	}
	return a
}

func FindCycleLength(start string, directions []bool, theMap map[string][]string) int {
	found := false
	foundFirstZ := false
	index := 0
	steps := 0
	location := start
	var firstZ string
	for {
		left := directions[index]
		
		location = GoNext(location, left, theMap)
		steps++	
	
		if !foundFirstZ {
			if location[2] == 'Z' {
				firstZ = location
			}
		}
		if strings.EqualFold(location, firstZ) {
			if steps % len(directions) == 0 {
			found = true
			}
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
	return steps
}

func GetGhostStarts(theMap map[string][]string) (ghostStarts []string) {
	for possibleStart := range theMap {
		if possibleStart[2] == 'A' {
			ghostStarts = append(ghostStarts, possibleStart)
		}
	}
	return ghostStarts
}

func CheckGhostFinish(locations []string) bool {
	for _, location := range locations {
		if location[2] != 'Z' {
			return false
		}
	}
	return true
}

func GoNext(start string, left bool, theMap map[string][]string) string {
	options := theMap[start]
	if left {
		return options[0]
	}
	return options[1]
}
