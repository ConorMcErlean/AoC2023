package main

import (
	"adventOfCode23/cmd/common"
	. "adventOfCode23/cmd/day3/fileReader"
	"fmt"
)

func main() {
	var lines = common.ReadFileFromArgs()
	var partNumbers []PartNumber
	var symbols []Coordinate
	var validPartNumbers []int
	for index, line := range lines {
		partNumbers, symbols = ParseLine(line, index, partNumbers, symbols)
		//fmt.Print("\n ---LINE ---")
		//fmt.Printf("\nPart Numbers %v", partNumbers)
		//fmt.Printf("\nSymbols %v", symbols)
	}

	for _, partNumber := range partNumbers {
		// Go through each number and check for hits
		validPartNumbers = checkForNearbySymbol(partNumber, symbols, validPartNumbers)
	}

	fmt.Printf("\nAll PartNumbers %v", validPartNumbers)
	var sum = 0
	for _, partNumber := range validPartNumbers {
		sum += partNumber
	}
	fmt.Printf("\nSum of PartNumbers = %v\n", sum)
}

func checkForNearbySymbol( part PartNumber, symbols []Coordinate, partNumbers []int) []int {
	var minX = part.StartCoordinate.X -1
	var minY = part.StartCoordinate.Y -1
	var maxX = part.EndCoordinate.X +1
	var maxY = part.EndCoordinate.Y +1

//	fmt.Printf("\nAnything in X %v - %v and Y %v - %v should be good", minX, maxX, minY, maxY)
	for _, symbol := range symbols {
		if symbol.X >= minX && symbol.X <= maxX {
			if symbol.Y >= minY && symbol.Y <= maxY {
				partNumbers = append(partNumbers,  part.Number)
			}
		}
	}
	return partNumbers
}

