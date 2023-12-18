package nestwalls

import (
	. "adventOfCode23/cmd/day10/pipes"
	"fmt"

)

type ReadRange struct {
	High int
	Low int
	Left int
	Right int
}

func BuildWalls(node PipeNode, diagram [][]PipeNode) ( []PipeNode, bool, ReadRange, [][]PipeNode ) {
	var wallsUp, wallsDown []PipeNode
	connectingWalls := false
	var readRange, readRange2 ReadRange
	// move left until not a .
	reverse := false
	for {
		if node.Char != '.' {
			break
		}
		if node.Location.Y == 0 {
			reverse = true 
		}
		if !reverse {
			node = diagram[node.Location.X][node.Location.Y-1]
		} else {
			node = diagram[node.Location.X][node.Location.Y+1]
		}
	}

	wallsUp, connectingWalls, readRange, diagram = buildWallInClockwiseDirection(node, diagram)
	if connectingWalls {
		return wallsUp, connectingWalls, readRange, diagram
	}

	wallsDown, connectingWalls, readRange2, diagram = buildWallInAntiClockwiseDirection(node, diagram)

	walls := append(wallsUp, wallsDown...)
	left, right, up, down := readRange.Left, readRange.Right, readRange.High, readRange.Low
	if readRange2.Left < left {
		left = readRange2.Left
	}
	if readRange2.High < up {
		up = readRange2.High
	}
	if readRange2.Right > right {
		right = readRange2.Right
	}
	if readRange2.Low> down {
		down = readRange2.Low
	}

	return walls, false, ReadRange{High: up, Low: down, Left: left, Right: right}, diagram
}

func buildWallInClockwiseDirection(node PipeNode, diagram [][]PipeNode) (walls []PipeNode, wallConnect bool, readRange ReadRange, returnDiagram [][]PipeNode ){
	start := Location{X: node.Location.X, Y: node.Location.Y }
	validWall, started, safe := false, false, true
	high, low := node.Location.X, node.Location.X
	left, right := node.Location.Y, node.Location.Y
	for {

		fmt.Printf("\n Checking %v", node.Location)

		// Check Reading range
		if node.Location.X > low {
			low = node.Location.X
		} else if node.Location.X < high {
			high = node.Location.X
		}

		if node.Location.Y > right {
			right = node.Location.Y
		} else if node.Location.Y < left {
			left = node.Location.Y
		}


		if !safe {
			fmt.Println("Am I out of bounds?")
			break
		}

		// Check back at start	
		if node.Location.X == start.X && node.Location.Y == start.Y {
			if started {
				fmt.Println("Done")
				validWall = true
				break
			} else {
				started = true
			}
		}

		diagram[node.Location.X][node.Location.Y].Checked = true


		// Corner
		node, validWall = topLeftCorner(node, diagram)
		if (validWall) {
			fmt.Print(" topLeft")
			var nextNode PipeNode
			nextNode, walls, safe = appendAndMoveRight(node, walls, diagram)
			if justBeenThere(nextNode, walls) {
				nextNode, safe = safeDown(node, diagram)
			}

			node = nextNode
			continue
		}
		// Corner
		node, validWall = topRightCorner(node, diagram)
		if (validWall) {
			fmt.Print(" topRight")
			var nextNode PipeNode

			nextNode, walls, safe = appendAndMoveDown(node, walls, diagram)
			if justBeenThere(nextNode, walls) {
				nextNode, safe = safeLeft(node, diagram)
			}
			node = nextNode

			continue

		}
		// Corner
		node, validWall = bottomRightCorner(node, diagram)
		if (validWall) {
			fmt.Print(" bottomRight")
			var nextNode PipeNode

			nextNode, walls, safe = appendAndMoveLeft(node, walls, diagram)
			if justBeenThere(nextNode, walls) {
				nextNode, safe = safeUp(node, diagram)
			}
			node = nextNode

			continue

		}
		// corner
		node, validWall = bottomLeftCorner(node, diagram)
		if (validWall) {
			fmt.Print(" bottomLeft")
			var nextNode PipeNode


			nextNode, walls, safe = appendAndMoveUp(node, walls, diagram)
			if justBeenThere(nextNode, walls) {
				nextNode, safe = safeRight(node, diagram)
			}
			node = nextNode
			continue
		}


		// Left
		node, validWall = leftWall(node, diagram)
		if (validWall) {
			fmt.Print(" Left")

			node, walls, safe = appendAndMoveUp(node, walls, diagram)
			continue
		}
		// top
		node, validWall = upWall(node, diagram)
		if (validWall) {
			fmt.Print(" Top")

			node, walls, safe = appendAndMoveRight(node, walls, diagram)
			continue

		}
		// Right
		node, validWall = rightWall(node, diagram)
		if (validWall) {
			fmt.Print(" Right")

			node, walls, safe = appendAndMoveDown(node, walls, diagram)
			continue
		}
		// bottom
		node, validWall = bottomWall(node, diagram)
		if (validWall) {
			fmt.Print(" bottom")

			node, walls, safe = appendAndMoveLeft(node, walls, diagram)
			continue
		}
		fmt.Print("Nothing?")



	}

	readRange = ReadRange { High: high+1, Low: low-1, Left: left+1, Right: right-1}

	return walls, validWall, readRange, diagram
}

// todo: Adjust directions
func buildWallInAntiClockwiseDirection(node PipeNode, diagram [][]PipeNode) (walls []PipeNode, wallConnect bool, readRange ReadRange, returnDiagram [][]PipeNode ){
	start := Location{X: node.Location.X, Y: node.Location.Y }
	high, low := node.Location.X, node.Location.X
	left, right := node.Location.Y, node.Location.Y
	validWall, started, safe := false, false, true
	for {
		// Check Reading range
		if node.Location.X > low {
			low = node.Location.X
		} else if node.Location.X < high {
			high = node.Location.X
		}

		if node.Location.Y > right {
			right = node.Location.Y
		} else if node.Location.Y < left {
			left = node.Location.Y
		}

		if !safe {
			fmt.Println("Am I out of bounds?")
			break
		}

		// Check back at start	
		if node.Location.X == start.X && node.Location.Y == start.Y {
			if started {
				fmt.Println("Done")
				validWall = true
				break
			} else {
				started = true
			}
		}

		diagram[node.Location.X][node.Location.Y].Checked = true



		// corner
		node, validWall = bottomLeftCorner(node, diagram)
		if (validWall) {
			fmt.Print(" bottomLeft")
			node, walls, safe = appendAndMoveRight(node, walls, diagram)
			continue
		}
		// Corner
		node, validWall = bottomRightCorner(node, diagram)
		if (validWall) {
			fmt.Print(" bottomRight")
			node, walls, safe = appendAndMoveUp(node, walls, diagram)
			continue

		}
		// Corner
		node, validWall = topRightCorner(node, diagram)
		if (validWall) {
			fmt.Print(" topRight")
			node, walls, safe = appendAndMoveLeft(node, walls, diagram)
			continue

		}
		// Corner
		node, validWall = topLeftCorner(node, diagram)
		if (validWall) {
			fmt.Print(" topLeft")
			node, walls, safe = appendAndMoveDown(node, walls, diagram)
			continue
		}

		// Left
		node, validWall = leftWall(node, diagram)
		if (validWall) {
			fmt.Print(" Left")

			node, walls, safe = appendAndMoveDown(node, walls, diagram)
			continue
		}

		// bottom
		node, validWall = bottomWall(node, diagram)
		if (validWall) {
			fmt.Print(" bottom")

			node, walls, safe = appendAndMoveRight(node, walls, diagram)
			continue

		}
		// Right
		node, validWall = rightWall(node, diagram)
		if (validWall) {
			fmt.Print(" right")

			node, walls, safe = appendAndMoveUp(node, walls, diagram)
			continue
		}
		// top
		node, validWall = upWall(node, diagram)
		if (validWall) {
			fmt.Print(" up")

			node, walls, safe = appendAndMoveLeft(node, walls, diagram)
			continue

		}


	
	}

	readRange = ReadRange { High: high+1, Low: low-1, Left: left+1, Right: right-1}


	return walls, validWall, readRange, diagram
}

func leftWall(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	right, safe := safeRight(node, diagram)
	if safe {
		if right.Char == '.' {
			return node, true
		}
	}
	return node, false
}

func rightWall(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	left, safe := safeLeft(node, diagram)
	if safe {
		if left.Char == '.' {
			return node, true
		}
	}
	return node, false
}

func topLeftCorner(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	x, y := node.GetLocation()
	right, safe1 := safeRight(node, diagram)
	down, safe2 := safeDown(node, diagram)
	if !safe1 || !safe2 {
		return node, false
	}
	diag := diagram[x+1][y+1]
	cornerPoint := down.Char != '.' && right.Char != '.'

	if diag.Char == '.' && cornerPoint {
		return node, true
	}
	return node, false
}


func topRightCorner(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	x, y := node.GetLocation()
	left, safe1 := safeLeft(node, diagram)
	down, safe2 := safeDown(node, diagram)
	if !safe1 || !safe2 {
		return node, false
	}
	diag := diagram[x+1][y-1]
	cornerPoint := down.Char != '.' && left.Char != '.'

	if diag.Char == '.' && cornerPoint {
		return node, true
	}
	return node, false
}

func upWall(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	down, safe := safeDown(node, diagram)
	if safe {
		if down.Char == '.' {
			return node, true
		}
	}
	return node, false
}


func bottomLeftCorner(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	x, y := node.GetLocation()
	right, safe1 := safeRight(node, diagram)
	up, safe2 := safeUp(node, diagram)
	if !safe1 || !safe2 {
		return node, false
	}

	diag := diagram[x-1][y+1]
	cornerPoint := right.Char != '.' && up.Char != '.'

	if diag.Char == '.' && cornerPoint {
		return node, true
	}
	return node, false
}

func bottomRightCorner(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	x, y := node.GetLocation()
	left, safe1 := safeLeft(node, diagram)
	up, safe2 := safeUp(node, diagram)
	if !safe1 || !safe2 {
		return node, false
	}
	diag := diagram[x-1][y-1]
	cornerPoint := up.Char != '.' && left.Char != '.'

	if diag.Char == '.' && cornerPoint {
		return node, true
	}	
	return node, false
}

func bottomWall(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	up, safe := safeUp(node, diagram)
	if safe {
		if up.Char == '.' {
			return node, true
		}
	}
	return node, false
}

func appendAndMoveUp(node PipeNode, walls []PipeNode, diagram [][]PipeNode) (PipeNode, []PipeNode, bool) {
	walls = append(walls, node)
	var safe bool
	node, safe = safeUp(node, diagram)	
	return node, walls, safe
}

func safeUp(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	if node.Location.X != 0 {
		return diagram[node.Location.X-1][node.Location.Y], true
	}
	return node, false
}

func appendAndMoveDown(node PipeNode, walls []PipeNode, diagram [][]PipeNode) (PipeNode, []PipeNode, bool) {
	walls = append(walls, node)
	var safe bool
	node, safe = safeDown(node, diagram)
	return node, walls, safe
}

func safeDown(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	if node.Location.X != len(diagram) -1 {
		return diagram[node.Location.X+1][node.Location.Y], true
	}
	return node, false
}

func appendAndMoveRight(node PipeNode, walls []PipeNode, diagram [][]PipeNode) (PipeNode, []PipeNode, bool) {
	walls = append(walls, node)
	var safe bool
	node, safe = safeRight(node, diagram)
	return node, walls, safe
}

func safeRight(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	if node.Location.Y != len(diagram[0]) -1 {
		return diagram[node.Location.X][node.Location.Y+1], true
	}
	return node, false
}

func appendAndMoveLeft(node PipeNode, walls []PipeNode, diagram [][]PipeNode) (PipeNode, []PipeNode, bool) {
	walls = append(walls, node)
	var safe bool
	node, safe = safeLeft(node, diagram)
	return node, walls, safe
}

func safeLeft(node PipeNode, diagram [][]PipeNode) (PipeNode, bool) {
	if node.Location.Y != 0 {
		return diagram[node.Location.X][node.Location.Y-1], true
	}
	return node, false
}

func justBeenThere(nextNode PipeNode, walls []PipeNode) bool {
	if len(walls) == 1 {
		return false
	}
	return nextNode == walls[len(walls) -2]
	}
