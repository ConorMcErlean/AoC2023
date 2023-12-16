package nests

import (
	. "adventOfCode23/cmd/day10/pipes"
	"fmt"
	"strings"
	. "adventOfCode23/cmd/day10/walls"
)
func FindEnclosed(diagram [][]rune) int {
	enclosed := 0

	alreadyChecked := make([][]bool, len(diagram))
	for index := range alreadyChecked {
		row := make([]bool, len(diagram[0]) )
		for i := range row {
			row[i] = false
		}
		alreadyChecked[index] = row
	}
	
	x,y, enclosed := 0, 0, 0
	var enclosedLocs []Location


	for {
		
		if !alreadyChecked[x][y] {
			var currentLocs []Location
			count := 0
			location := Location {X: x, Y: y }
			count, alreadyChecked, currentLocs = Check(location, diagram, alreadyChecked)
			enclosed += count
			if count > 0 {
				enclosedLocs = append(enclosedLocs, currentLocs... ) 
			}
		}

		if x == len(alreadyChecked) -1 && y == len(alreadyChecked[0])-1 {
			break
		}

		if y < len(alreadyChecked[0]) -1 {
			y++
			continue
		} else {
			y = 0
			x++
		}
	}

	fmt.Printf("\nEnclosed Locations: %v\n", enclosedLocs)

	return enclosed
}

func Check(
	startlocation Location, 
	diagram [][]rune, 
	checked [][]bool )(int, [][]bool, []Location) {
	count :=0
	location := startlocation
	var currentLocs, empty []Location

	char := diagram[startlocation.X][startlocation.Y]
	if char == 'J' || char == 'F' || char == 'L' || char == '7' || char == '-' || char == '|'{
		checked[startlocation.X][startlocation.Y] = true
		return 0, checked, empty
	}


	for {
		checked[location.X][location.Y]=true

		if (FoundOuterWall(location, diagram)){
			return 0, checked, empty
		}
		if LowerWall(location, diagram) || LeftWall(location, diagram) || RightWall(location, diagram){
			return count, checked, currentLocs
		}
		
		count++
		currentLocs = append(currentLocs, location)

		count, checked, currentLocs = CheckDir("left", Location{X: location.X, Y: location.Y -1}, diagram, checked, count, currentLocs)
		count, checked, currentLocs = CheckDir("right", Location{X: location.X, Y: location.Y +1}, diagram, checked, count, currentLocs)
		if count == 0 {
			return 0, checked, empty
		}
		location = Location{X: location.X +1, Y: location.Y}
	}
}

func CheckDir(
	direction string, 
	location Location, 
	diagram [][]rune, 
	checked [][]bool, 
	count int, 
	currentLocs []Location ) (int, [][]bool, []Location) {
	// Incase you are sent an out of index location
	if location.X < 0 || location.X > len(checked) -1 || location.Y < 0 || location.Y > len(checked[0]){
		return 0, checked, currentLocs
	}

	left := strings.EqualFold("left", direction)

	count, checked, currentLocs = HorizontalPipe(direction, location, diagram, checked, count, currentLocs)

	count, checked, currentLocs = VerticalPipe(direction, location, diagram, checked, count, currentLocs)


	if left || LeftWall(location, diagram) {
		return count, checked, currentLocs	
	} else if RightWall(location, diagram) {
		return count, checked, currentLocs	
	}

	if (FoundOuterWall(location, diagram)){
		return 0, checked, currentLocs
	}

	// May need to double check this logic
	if checked[location.X][location.Y] {
		return count, checked, currentLocs
	} else {
		checked[location.X][location.Y] = true
	}
	count++
	currentLocs = append(currentLocs, location)

	if left {
		count, checked, currentLocs = CheckDir(direction, Location{X: location.X, Y: location.Y-1 }, diagram, checked, count, currentLocs)
	} else {
		count, checked, currentLocs = CheckDir(direction, Location{X: location.X, Y: location.Y+1 }, diagram, checked, count, currentLocs)
	}

	return count, checked, currentLocs
}


func HorizontalPipe(
	direction string, 
	location Location, 
	diagram [][]rune, 
	checked [][]bool, 
	count int, 
	currentLocs []Location ) (int, [][]bool, []Location) {
	char := diagram[location.X][location.Y]
	var nextDir Location
	if strings.EqualFold("left", direction) {
		nextDir = Location{X: location.X, Y: location.Y -1}
	} else {
		nextDir = Location{X: location.X, Y: location.Y+1 }
	}

	if char == 'F' || char == '7' || char == '-' {
		if location.X == 0 {
			return count, checked, currentLocs
		}
		lowerChar := diagram[location.X+1][location.Y]
		// Check if just wall, or Pipe
		if lowerChar == 'J' || lowerChar == 'L' || lowerChar == '-' {
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}

	if char == 'J' ||char == 'L' || char == '-' {
		if location.X == len(diagram) -1 {
			return count, checked, currentLocs
		}
		upperChar := diagram[location.X-1][location.Y]
		if upperChar == 'F' || upperChar == '7' || upperChar == '-' {
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}
	return count, checked, currentLocs
}

func VerticalPipe(
	direction string, 
	location Location, 
	diagram [][]rune, 
	checked [][]bool, 
	count int, 
	currentLocs []Location ) (int, [][]bool, []Location) {

	char := diagram[location.X][location.Y]
	var nextDir Location
	nextDir = Location{X: location.X-1, Y: location.Y }

	if char == 'F' || char == 'L' || char == '|' {
		if location.Y == 0 {
			return count, checked, currentLocs
		}
		leftChar := diagram[location.X][location.Y-1]
		// Check if just wall, or Pipe
		if leftChar == 'J' || leftChar == '7' || leftChar == '|' {
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}

	if char == 'J' ||char == '7' || char == '|' {
		if location.Y == len(diagram[0]) -1 {
			return count, checked, currentLocs
		}
		rightChar := diagram[location.X][location.Y+1]
		if rightChar == 'F' || rightChar == 'L' || rightChar == '|' {
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}
	return count, checked, currentLocs
}
