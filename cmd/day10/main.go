package main

import (
	"adventOfCode23/cmd/day10/nests"
	"adventOfCode23/cmd/day10/pipes"
	"fmt"
)

func main() {
	pipeDiagram, start := pipes.ReadPipeDiagram()
	route1, route2 := pipes.GetRoutes(pipeDiagram, start)
	FindMiddle(route1, route2)
	pipeOutline := pipes.PrintPipes(pipeDiagram, route1)
	enclosed := nests.FindEnclosed(pipeOutline)
	fmt.Printf("\n Enclosed in above %v\n", enclosed)



}

func FindMiddle(route1 []pipes.Location, route2[]pipes.Location) {
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
