package filereader

import (
	"unicode"
	"fmt"
	"strconv"
	"log"
)

func ParseLine(
	line string, 
	lineIndex int,
	partNumbers []PartNumber,
	symbols []Coordinate,
	gears []Gear,
) ([]PartNumber, []Coordinate, []Gear) {
	var startCoOrd, endCoOrd, number = Reset()
	var lineLength = len(line)

	for index, char := range line {
		switch {
		case unicode.IsNumber(char):
			// Start building our part number
			if IsIntroCoOrd(startCoOrd) {
				startCoOrd = Coordinate{ X: index, Y: lineIndex}
			}
			number += string(char)
			endCoOrd = Coordinate{ X: index, Y: lineIndex}
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			// Stop Building our number
			// Check if anythings been set, Make a Part Number then Reset
			partNumbers =  StopBuildingAndAdd(startCoOrd, endCoOrd, number, partNumbers)
			startCoOrd, endCoOrd, number = Reset() 
			if char != '.' {
				location := Coordinate{ X: index, Y: lineIndex}
				symbols = append(symbols, location) 
				if char == '*' {
					gears = append(gears, Gear { Location : location } )
				}
			}
		default:
			fmt.Println(string(char), "Is neither a number or Symbol")
		}
		
		// Irrespective of the above result, if its the end of the line we need to close out any building partNumbers
		if index == (lineLength -1) {
			// End of a Line also stop building our number
			partNumbers =  StopBuildingAndAdd(startCoOrd, endCoOrd, number, partNumbers)
			startCoOrd, endCoOrd, number = Reset() 
		}
	}
	return partNumbers, symbols, gears
}

func StopBuildingAndAdd(start Coordinate, end Coordinate, number string, partNumbers []PartNumber) []PartNumber {
	if !IsIntroCoOrd(start) {
		partNum := PartNumber {
			Number : StringToInt(number),
			StartCoordinate : start,
			EndCoordinate : end,
		}
		partNumbers = append(partNumbers, partNum)
	}
	return partNumbers
}

func IsIntroCoOrd(coordinate Coordinate) bool {
	return (coordinate.X == -1) && (coordinate.Y == -1)
}

func Reset() (Coordinate, Coordinate, string) {
	return Coordinate{ X: -1, Y: -1}, Coordinate{ X: -1, Y: -1}, ""
}

func StringToInt(number string) int {
	var num, err = strconv.Atoi(number)
	if (err != nil) {
		log.Fatalln("Parsing Error!", err)
	}
	return num
}

type PartNumber struct {
	Number int
	StartCoordinate Coordinate
	EndCoordinate Coordinate
}

type Gear struct {
	Location Coordinate
	PartNumbers []PartNumber
}

type Coordinate struct {
	X int
	Y int
}
