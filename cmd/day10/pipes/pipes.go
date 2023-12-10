package pipes

import (
	"adventOfCode23/cmd/common"
)

func ReadPipeDiagram() ( [][]rune, Location ) {
	file := common.ReadFileFromArgs()
	pipes := make([][]rune, len(file))
	var start Location

	for x, line := range file {
		pipes[x] = make([]rune, len(line))
		for y, char := range line {
			pipes[x][y] = char
			if char == 'S' {
				start = Location{ X:x, Y:y }
			}
		}
	}
	return pipes, start
}

func GetRoutes(diagram [][]rune, start Location)([]Location, []Location){
	start1, start2 := FindValid(diagram, start)
	way1 := BuildRoute(diagram, start, start1)
	way2 := BuildRoute(diagram, start, start2)
	return way1, way2
}

func BuildRoute(diagram [][]rune, start Location, first Location) (locations []Location){
	location := first
	last := start
	locations = append(locations, start)
	for {
		char := diagram[location.X][location.Y]
		next1, next2 := CheckValidPipes(location, char)
		locations = append(locations, location)
		if next1.X == last.X && next1.Y == last.Y {
			last = location
			location = next2
		} else {
			last = location
			location = next1
		}

		if location.X == -1 {
			break
		}
	}
	return locations
}

func CheckValidPipes(location Location, char rune) (Location, Location) {
	x := location.X
	y := location.Y
	var loc1, loc2 Location
	switch(char) {
	case '|' : 
		loc1 = Location{X: x-1, Y: y }
		loc2 = Location{X: x+1, Y: y }
	case '-' :
		loc1 = Location{X: x, Y: y-1 }
		loc2 = Location{X: x, Y: y+1 }
	case 'L' :
		loc1 = Location{X: x-1, Y: y }
		loc2 = Location{X: x, Y: y+1 }
	case 'J' :
		loc1 = Location{X: x-1, Y: y }
		loc2 = Location{X: x, Y: y-1 }
	case '7' :
		loc1 = Location{X: x+1, Y: y }
		loc2 = Location{X: x, Y: y-1 }
	case 'F' :
		loc1 = Location{X: x+1, Y: y }
		loc2 = Location{X: x, Y: y+1 }
	case 'S' :
		// End
		loc1 = Location{X: -1, Y: -1}
		loc2 = Location{X: -1, Y: -1}
	}
	return loc1, loc2
}

func FindValid(diagram [][]rune, start Location) ( Location, Location ){
	x := start.X
	y := start.Y
	var valid []Location
	up, down, left, right := 'N', 'N', 'N', 'N'	
	if x != 0 {
		up = diagram[x-1][y]
	}

	if x != len(diagram)-1 {
		down = diagram[x+1][y]
	}
	if y != 0 {
		left = diagram[x][y-1]
	}
	if y != len(diagram[0])-1 {
		right = diagram[x][y+1]
	}	


	if up == '|' || up == '7' || up == 'F' {
		valid = append(valid, Location{X: x-1, Y: y})
	}
	if down == '|' || down == 'L' || down == 'J' {
		valid = append(valid, Location{X: x+1, Y: y})
	}
	if left == '-' || left == 'L' || left == 'F' {
		valid = append(valid, Location{X: x, Y: y-1})
	}
	if right == '-' || right == 'J' || right == '7' {
		valid = append(valid, Location{X: x, Y: y+1})
	}
	return valid[0], valid[1]
}

type Location struct {
	X int
	Y int
}
