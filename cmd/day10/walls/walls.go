package walls

import(
	. "adventOfCode23/cmd/day10/pipes"
)
func FoundOuterWall(location Location, diagram [][]rune) bool {
	if location.X ==0 || location.X == len(diagram)-1 {
		return true
	}
	if location.Y == 0 || location.Y == len(diagram[0])-1 {
		return true
	}
	return false
}

func LeftWall(location Location, diagram [][]rune) bool {
	char := diagram[location.X][location.Y]

	if char == 'J' || char == '7' || char == '|' {
		return true
	}
	return false
}

func RightWall(location Location, diagram [][]rune) bool {
	char := diagram[location.X][location.Y]

	if char == 'F' || char == 'L' || char == '|' {
		return true
	}
	return false
}

func UpperWall(location Location, diagram [][]rune) bool {
	char := diagram[location.X][location.Y]

	if char == 'J' || char == 'L' || char == '-' {
		return true
	}
	return false
}

func LowerWall(location Location, diagram [][]rune) bool {
	char := diagram[location.X][location.Y]

	if char == 'F' || char == '7' || char == '-' {
		return true
	}
	return false
}

