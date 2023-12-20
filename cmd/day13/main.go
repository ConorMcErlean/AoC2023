package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"slices"
	"strings"
)

func main() {
	file := common.ReadFileFromArgsWithEmptyLines()
	maps := parseToMaps(file)
	total :=0
	for i, mapp := range maps {
		count := 0
				
		fmt.Printf("\n Checking map %v of %v", i+1, len(maps))
		mirrorPointHor, foundH := findHorizontalReflection(mapp)
		mirrorPointVert, foundV := findVerticalReflection(mapp)

		if foundH {
			count = mirrorPointHor
		}

		if foundV {
			count += (100* mirrorPointVert)
		}
		total += count
	}
	fmt.Printf("\n\nTotal %v\n", total)
}

func parseToMaps(input []string) [][]string {
	var maps [][]string 
	var currentMap []string

	for i, line := range input {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			currentMap = append(currentMap, line)
		} else {
			maps = append(maps, currentMap)
			currentMap = make([]string, 0)

		}
		if i == len(input)-1 && len(currentMap) != 0 {
			maps = append(maps, currentMap)
		}
	}
	return maps
}

func findHorizontalReflection(input []string) (int, bool) {
	var horizontal, lastHorizontal []rune
	var possibleReflections []int 

	for i := range input[0] {
		horizontal = createColumn(i, input)	
		if i == 0 {
			lastHorizontal = horizontal
			continue 
		}
		if slices.Equal(horizontal, lastHorizontal) {
			possibleReflections = append(possibleReflections, i)
		}
		lastHorizontal = horizontal
	}

	for _, reflection := range possibleReflections {
		left, right := reflection -1, reflection
		for {
			if !(left >= 0) || !(right < len(input[0])) {
				return reflection, true
			}


			leftColumn, rightColumn := createColumn(left, input), createColumn(right, input)
		//	fmt.Printf("\nComparing\n%v\n%v", leftColumn, rightColumn)
			if !slices.Equal(leftColumn, rightColumn ) {
				break
			}

			left--
			right++

				}
	}
	return -1, false
}

func findVerticalReflection(input []string) (int, bool) {
	var lastline string
	var possibleReflections []int 

	for index, line := range input {
		if index == 0 {
			lastline = line
			continue
		}
		if strings.EqualFold(lastline, line) {
			possibleReflections = append(possibleReflections, index)
		}
		lastline = line
	}

	for _, reflection := range possibleReflections {
		up, down := reflection -1, reflection
		for {
			if !(up >= 0) || !(down < len(input) ) {
				return reflection, true
			}

			upString, downString := input[up], input[down]
		//	fmt.Printf("\nComparing\n%v\n%v", upString, downString)
			if !strings.EqualFold(upString, downString) {
				break
			}
			down++
			up --
		}
	}
	return -1, false
}

func createColumn(index int, input []string) []rune {
	column := make([]rune, len(input))

	for row := range input { 
		column[row] = rune(input[row][index])
	}	

	return column
}
