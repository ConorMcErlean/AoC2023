package nests

import (
	nestwalls "adventOfCode23/cmd/day10/nestWalls"
	. "adventOfCode23/cmd/day10/pipes"
	. "adventOfCode23/cmd/day10/walls"
	"fmt"
	"strings"
)

type Nest struct {
	Wall []PipeNode
	Contents int
	Valid bool
}

//func (node *PipeNode) AssociateNest(nest Nest) {
//	node.AssociatedNests = append(node.AssociatedNests, nest)
//}

func BuildNest(node PipeNode, diagram [][]PipeNode) (Nest, [][]PipeNode) {
	// Builds Walls
	wall, valid, readRange, diagram := nestwalls.BuildWalls(node, diagram)
	contents, diagram := calculateContents(wall, readRange, diagram)
	if !valid {
		contents = 0
	}
	return Nest{ Wall: wall, Contents: contents, Valid: valid }, diagram
}

func (nest Nest) checkTunnels(diagram [][]PipeNode) {
	// todo

}

// Some flaw with this method...
func calculateContents(walls []PipeNode, readRange nestwalls.ReadRange, diagram [][]PipeNode) (int, [][]PipeNode) {
	// should be safe access & doesnt handle open nests well
	reading := checkPointInWalls(diagram[readRange.High-1][readRange.Left-1], walls)
	count := 0
	ignoreNextWall := false
	
	fmt.Printf("Read range high %v, low %v, left %v right%v", readRange.High, readRange.Low, readRange.Left, readRange.Right)
	for i := readRange.High; i <= readRange.Low; i++ {
		for j := readRange.Left; j < readRange.Right; j++ {
			if !reading {
				fmt.Printf("\n Im not reading for node %v because high %v & left %v were not in the wall", diagram[i][j], readRange.High, readRange.Left)
			}
			if reading {
				fmt.Printf("\nI am Reading %v:%v\n", i, j)
				point := diagram[i][j]
				if point.Char == '.' {
					fmt.Println("I am Counting ")
					count++
				} else if checkPointInWalls(point, walls) {
					// Need to handle a chain of wall Locations
					if !ignoreNextWall {
						reading = !reading
						ignoreNextWall = true
					} else if checkPointInWalls(diagram[i][j+1], walls) {
						continue
					} else {
						reading = !reading
						ignoreNextWall = false
					}
				}

				diagram[point.Location.X][point.Location.Y].Checked = true
				fmt.Printf("Should be true %v:%v %v count %v", point.Location.X, point.Location.Y, diagram[point.Location.X][point.Location.Y].Checked, count)
			}
		}
	}
	return count, diagram
}


func buildNodeMap(diagram [][]rune) [][]PipeNode {
	nodeMap := make([][]PipeNode, len(diagram))
	for i := range diagram {
		nodeMap[i] = make([]PipeNode, len(diagram[i]))
		for j, char := range diagram[i] {
			nodeMap[i][j] = PipeNode{
				Location: Location{X: i, Y: j},
				Checked: false,
				Char: char,
			}
		}
	}
	return nodeMap
}

func checkPointInWalls(node PipeNode, walls []PipeNode) bool {
	for _, wallNode := range walls {
		if wallNode.Location.X == node.Location.X && wallNode.Location.Y == node.Location.Y {
			return true
		}
	}
	return false
}

func FindEnclosed(diagram [][]rune) int {
	nodeMap := buildNodeMap(diagram)

	var nests []Nest
	i, j := 0, 0

	for {
		node := nodeMap[i][j]
			if !node.Checked {
			if node.Char == '.' {
				var nest Nest
				nest, nodeMap = BuildNest(node, nodeMap)
				nests = append(nests, nest)
			}
		}
		fmt.Printf("\n has %v:%v been checked %v", i, j, node.Checked)


		// Loop Logic
		if j == len(nodeMap[i]) -1 {
			if i == len(nodeMap) -1 {
				break
			}
			j = 0;
			i++
		} else {
			j++
		}


	}

	total := 0

	for _, nest := range nests {
		if nest.Valid {
			fmt.Printf("\n A nest found with %v contents", nest.Contents )
			total += nest.Contents
		}
	}
	return total
}
// Assume below here is trash ------

///func FindEnclosed(diagram [][]rune) int {
//	enclosed := 0
///
//	alreadyChecked := make([][]bool, len(diagram))
//	for index := range alreadyChecked {
//		row := make([]bool, len(diagram[0]) )
//		for i := range row {
//			row[i] = false
//		}
//		alreadyChecked[index] = row
//	}
//	
//	x,y, enclosed := 0, 0, 0
//	var enclosedLocs []Location
//
//
//	for {
//		
//		if !alreadyChecked[x][y] {
//			var currentLocs []Location
//			count := 0
//			location := Location {X: x, Y: y }
//			count, alreadyChecked, currentLocs = Check(location, diagram, alreadyChecked)
//			enclosed += count
//			if count > 0 {
//				enclosedLocs = append(enclosedLocs, currentLocs... ) 
//			}
//		}
//
//		if x == len(alreadyChecked) -1 && y == len(alreadyChecked[0])-1 {
//			break
//		}
//
//		if y < len(alreadyChecked[0]) -1 {
//			y++
//			continue
//		} else {
//			y = 0
//			x++
//		}
//	}
//
//	fmt.Printf("\nEnclosed Locations: %v\n", enclosedLocs)
//
//	return enclosed
//}

func Check(
	startlocation Location, 
	diagram [][]rune, 
	checked [][]bool )(int, [][]bool, []Location) {
	count := 0
	var currentLocs, empty []Location

	char := diagram[startlocation.X][startlocation.Y]
	if char == 'J' || char == 'F' || char == 'L' || char == '7' || char == '-' || char == '|'{
		checked[startlocation.X][startlocation.Y] = true
		return 0, checked, empty
	}

	//fmt.Printf("\n Checking %v", startlocation)

	count, checked, currentLocs = CheckDown(startlocation, diagram, checked, count, currentLocs)
	//fmt.Printf("\n count After check %v", count)
	return count, checked, currentLocs
}

func CheckDown(
	location Location, 
	diagram [][]rune, 
	checked [][]bool, 
	count int, 
	currentLocs []Location ) (int, [][]bool, []Location)  {
	var empty []Location

	checked[location.X][location.Y]=true

	if (FoundOuterWall(location, diagram)){

		fmt.Println("Returning zero for outer wall")
		return 0, checked, empty
	}
	if LowerWall(location, diagram) || LeftWall(location, diagram) || RightWall(location, diagram){
		return count, checked, currentLocs
	}
		
	count++
	currentLocs = append(currentLocs, location)
	count, checked, currentLocs = CheckDir("left", Location{X: location.X, Y: location.Y -1}, diagram, checked, count, currentLocs)
	count, checked, currentLocs = CheckDir("right", Location{X: location.X, Y: location.Y +1}, diagram, checked, count, currentLocs)
	
	fmt.Printf("\n After Horizontal Check, count is %v", count)
	if count == 0 {
		fmt.Println("returning zero, because I got a zero")
		return 0, checked, empty
	}

	fmt.Printf("\n before going down a level, count is %v", count)

	count, checked, currentLocs = CheckDown(Location{X: location.X +1, Y: location.Y}, diagram, checked, count, currentLocs )
	return count, checked, currentLocs
	
}

func CheckDir(
	direction string, 
	location Location, 
	diagram [][]rune, 
	checked [][]bool, 
	count int, 
	currentLocs []Location ) (int, [][]bool, []Location) {
	fmt.Printf("\n On %v pass, value is %v", direction, count)
	// Incase you are sent an out of index location
	if location.X < 0 || location.X > len(checked) -1 || location.Y < 0 || location.Y > len(checked[0]){
		fmt.Println("Returning zero because out of bounds")
		return 0, checked, currentLocs
	}

	left := strings.EqualFold("left", direction)

	count, checked, currentLocs = HorizontalPipe(direction, location, diagram, checked, count, currentLocs)

	count, checked, currentLocs = VerticalPipe(location, diagram, checked, count, currentLocs)
	fmt.Printf("Count after pipes %v", count)
	
	// if on the sweep we find a path below, follow it
	if (location.X != len(checked) -1 ) && diagram[location.X+1][location.Y] == '.' {
		fmt.Println("going down!!")
		count, checked, currentLocs = CheckDown(Location{X: location.X +1, Y: location.Y}, diagram, checked, count, currentLocs)

	}


	if LeftWall(location, diagram) || RightWall(location, diagram) {
		fmt.Printf("\nRight wall returning %v", count)
		return count, checked, currentLocs	
	}
	// Need a way to go down
	if left {
		fmt.Println("Do we ever hit this line?")
	}
	if (FoundOuterWall(location, diagram)){
		fmt.Println("Returning zero for outer wall")
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
	fmt.Printf("\n HoriPipe starting with %v", count)

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
		upperChar := diagram[location.X-1][location.Y]
		// Check if just wall, or Pipe
		if upperChar == 'J' || upperChar == 'L' || upperChar == '-' {
			fmt.Println("PIPE Hor")
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}

	if char == 'J' ||char == 'L' || char == '-' {
		if location.X == len(diagram) -1 {
			return count, checked, currentLocs
		}
		lowerChar := diagram[location.X+1][location.Y]
		if lowerChar == 'F' || lowerChar == '7' || lowerChar == '-' {

			fmt.Println("PIPE Hor")
			count, checked, currentLocs = CheckDir(direction, nextDir, diagram, checked, count, currentLocs)
		}
	}

	fmt.Printf("\n HoriPipe returning %v", count)
	return count, checked, currentLocs
}

func VerticalPipe(
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

			fmt.Printf("\n in a pipe down going to %v", nextDir)
			count, checked, currentLocs = CheckDown(nextDir, diagram, checked, count, currentLocs)
		}
	}

	if char == 'J' ||char == '7' || char == '|' {
		if location.Y == len(diagram[0]) -1 {
			return count, checked, currentLocs
		}
		rightChar := diagram[location.X][location.Y+1]
		if rightChar == 'F' || rightChar == 'L' || rightChar == '|' {
			fmt.Printf("\n in a pipe down going to %v", nextDir)
			count, checked, currentLocs = CheckDown(nextDir, diagram, checked, count, currentLocs)
		}
	}
	fmt.Printf("\n VertPipe returning %v", count)
	return count, checked, currentLocs
}
