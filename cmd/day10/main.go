package main

import (
	"adventOfCode23/cmd/day10/pipes"
	"fmt"
)

func main() {
	pipeDiagram, start := pipes.ReadPipeDiagram()
	route1, route2 := pipes.GetRoutes(pipeDiagram, start)
	fmt.Println(route1)
	fmt.Println(route2)
}
