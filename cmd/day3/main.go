package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day3/fileReader"
	"fmt"
)

func main() {
	var partNumbers []PartNumber
	var symbols []Coordinate
	var gears []Gear
	var ratios []int
	var validPartNumbers []int

	// Parse Input File
	var lines = common.ReadFileFromArgs()
	for index, line := range lines {
		partNumbers, symbols, gears = ParseLine(line, index, partNumbers, symbols, gears)
	}

	// Find Part Numbers
	for _, partNumber := range partNumbers {
		// Go through each number and check for hits
		validPartNumbers = checkForNearbySymbol(partNumber, symbols, validPartNumbers)
	}

	// Find Gears, and thus Ratios
	ratios = checkForGears(partNumbers, gears)

	// fmt.Printf("\nAll PartNumbers %v", validPartNumbers)
	var sum, ratioSum = 0, 0
	
	// Summing up to complete AoC Questions
	for _, partNumber := range validPartNumbers {
		sum += partNumber
	}
	for _, ratios := range ratios {
		ratioSum += ratios
	}
	fmt.Printf("\nSum of PartNumbers = %v\nSum of Gear Ratios = %v\n", sum, ratioSum)

}

func checkForNearbySymbol( part PartNumber, symbols []Coordinate, partNumbers []int) []int {
	var minX = part.StartCoordinate.X -1
	var minY = part.StartCoordinate.Y -1
	var maxX = part.EndCoordinate.X +1
	var maxY = part.EndCoordinate.Y +1

	for _, symbol := range symbols {
		if symbol.X >= minX && symbol.X <= maxX {
			if symbol.Y >= minY && symbol.Y <= maxY {
				partNumbers = append(partNumbers,  part.Number)
			}
		}
	}
	return partNumbers
}

func checkForGears(parts []PartNumber, gears []Gear) []int {
	var gearLikes = make( map[int][]int )
	var ratios []int
	for gearIndex, gear := range gears {
		for _, part := range parts {
			var minX = part.StartCoordinate.X -1
			var minY = part.StartCoordinate.Y -1
			var maxX = part.EndCoordinate.X +1
			var maxY = part.EndCoordinate.Y +1

			if gear.Location.X >= minX && gear.Location.X <= maxX {
				if gear.Location.Y >= minY && gear.Location.Y <= maxY {
					fmt.Println("Gear Like! Part Numbers associated")
					gearLikes[gearIndex] = append(gearLikes[gearIndex], part.Number)
					fmt.Println(gearLikes)
				}
			}
		}
	}
	
	for _, partNumbers := range gearLikes {
		
		if len(partNumbers) == 2 {
			var ratio = partNumbers[0] * partNumbers[1]
			fmt.Printf("\nValid Gear near part numbers %v & %v with a ratio of %v\n", partNumbers[0], partNumbers[1], ratio)
			ratios = append(ratios, ratio)
		}
	}

	return ratios
}
