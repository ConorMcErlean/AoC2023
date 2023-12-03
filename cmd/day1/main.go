package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var numMap map[string]int
func main(){
	var calibrationFile = common.ReadFileFromArgs()
	initNumMap()
	parseAndPrintCalibration(calibrationFile)
}

func initNumMap(){
	numMap = make(map[string]int)
	numMap["zero"] = 0
	numMap["one"] = 1
	numMap["two"] = 2
	numMap["three"] = 3
	numMap["four"] = 4
	numMap["five"] = 5
	numMap["six"] = 6
	numMap["seven"] = 7
	numMap["eight"] = 8
	numMap["nine"] = 9
}

func parseAndPrintCalibration(calibrations []string) {
//	var calibrationValues []int
	
	total := 0
	for i := 0; i < len(calibrations); i++ {
		var calibrationRaw = calibrations[i] 
		if len(calibrationRaw) > 0 {
			var calibration = parseCalibration(calibrationRaw);
			fmt.Printf("%d ", calibration)
			//calibrationValues = append(calibrationValues, calibration)
			total += calibration
		}
	}
	// fmt.Printf("\nCalibration Values: %v \n", calibrationValues)
	fmt.Println("\nTotal = ", total)
}

func parseCalibration(input string) int {
	// Store numbers as index (key), number (value)
	var numbers = make(map[int]int)
	
	parseNumbersFromLine(input, numbers)
	parseWordsToNumbers(input, numbers)

	// By this stage numbers maps should contain all numbers and its index
	//for k, v := range numbers {
	//	log.Println("at index", k, "we have ", v)
	//}
	//log.Println("----")
	var first, last = checkMatchIndexes(numbers)
	// By this point we should have the true first & last
	var combo = strconv.Itoa(first) + strconv.Itoa(last)

	var cal, err = strconv.Atoi(combo)
	if (err != nil) {
		println("Parsing Error!", err)
	}
	return cal
}

func checkMatchIndexes(numbers map[int]int) (int, int) {
	// Initialise with some numbers that will always be replaced
	var firstNum, lastNum, lowestIndex, highestIndex = 0, 0, 999, -1
	// Now to compare indexes
	for index, number := range numbers {
		if index < lowestIndex {
			firstNum = number
			lowestIndex = index
		}
		if index > highestIndex {
			lastNum = number
			highestIndex = index
		}
	}
	return firstNum, lastNum
}

func parseNumbersFromLine(line string, numbers map[int]int ){
	// for part two we need to parse words as ints.
	// this means we will need to take substrings and check the starting byte.
	for index, char := range line {
		if (unicode.IsDigit(char)){
			// Add any numbers to the map and we can compare its index later
			numbers[index] = getNumFromChar(char)
		}
	}
}

func getNumFromChar(char rune) int {
	var number, err = strconv.Atoi(string(char))			
	if err != nil {
		log.Fatal("Couldnt properly parse a number from a rune")
	}
	return number
}


// for later remember a line could contain to numsAsWords so you need to remove as you find
func parseWordsToNumbers(line string, numbers map[int]int) {
	for numAsWord := range numMap {
		checkForNumber(line, numAsWord, numbers)
	}
}

func checkForNumber(line string, word string, numbers map[int]int) {
	if strings.Contains(line, word) {
		// get first and last instance of word in line
		// we dont care about a middle instance
		numbers[strings.Index(line, word)] = numMap[word]
		numbers[strings.LastIndex(line,word)] = numMap[word]	
	}
}
