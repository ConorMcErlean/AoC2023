package main

import (
	. "adventOfCode23/cmd/day10/pipes"
	"fmt"
	"slices"
)


// Part 2 Solution had me so lost I had to research & find a guide...
func main() {
	pipeDiagram, start := ReadPipeDiagram()
	route1, route2 := GetRoutes(pipeDiagram, start)
	FindMiddle(route1, route2)
	pipeOutline := PrintPipes(pipeDiagram, route1)
	outside := findEnclosed(pipeOutline)
	total := checkRemains(pipeOutline, route1, outside )

	fmt.Printf("\nTotal Inside %v", total)
}

func FindMiddle(route1 []Location, route2 []Location) {
	for i :=0; i < len(route1); i++ {
		loc1 := route1[i]
		loc2 := route2[i]
		
		if loc1.X == loc2.X && loc1.Y == loc2.Y {
			if (i != 0) && (i != len(route1) -1) {
				fmt.Printf("\nfound the middle %v, %v at %v", loc1, loc2, i)
			}
		}
	}
}

func checkRemains(pipes [][]rune, wall []Location, outside []Location) int {
	count := 0

	for r, row := range pipes {
		for c := range row {
			here := Location{X:r, Y:c }
			if !slices.Contains(wall, here) && !slices.Contains(outside, here) {
				count++
			}
		}
	}
	return count
}

func findEnclosed(pipes [][]rune) []Location {
	var outside []Location
	for r, row := range pipes {
		within := false
		up := false

		for c, char := range row {
			
			if char == '|' {
				within = !within
			} else if char == '-' {
				// Nothing?
			} else if char == 'L' || char == 'F' {
				if char == 'L' {
					up = true
				}
 
			} else if char == '7' || char == 'J' {
				if !up {
					if char != '7' {
						within = !within
					}
				} else {
					if char != 'J' {
						within = !within
					}
				}
				up = false
			}
			if !within {
				outside = append(outside, Location{X: r, Y: c})
			}
		}
	}
	return outside
}
