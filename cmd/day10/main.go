package main

import (
	"adventOfCode23/cmd/day10/pipes"
	"fmt"
)

func main() {
	pipeDiagram, start := pipes.ReadPipeDiagram()
	route1, route2 := pipes.GetRoutes(pipeDiagram, start)
//	fmt.Println(route1)
//	fmt.Println(route2)
	pipes.PrintPipes(pipeDiagram, route1)

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
