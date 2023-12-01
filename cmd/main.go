package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numMap map[string]int
func main(){
	var arguements = os.Args
	if len(arguements) < 2 {
		log.Fatal("We really need a file to do anything")
	}
	initNumMap()
	var calibrationFile = readFile(arguements[1])
	parseAndPrintCalibration(calibrationFile)
}

func readFile(filename string) []string {
	var fileLines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatalf("Something went wrong scanning: %v", err)
	}
	return fileLines
}

func parseAndPrintCalibration(calibrations []string) {
	var calibrationValues []int
	
	total := 0
	for i := 0; i < len(calibrations); i++ {
	        var calibration = parseCalibration(calibrations[i]);
		calibrationValues = append(calibrationValues, calibration)
		total += calibration
	}
	fmt.Printf("\nCalibration Values: %v \n", calibrationValues)
	fmt.Println("Total = ", total)
}

func parseCalibration(input string) int {
	var parsedFirstNumber, parsedLastNumber = parseNumbersFromLine(input)
	var parsedStringNumbers = parseLineToInts(input)
	// Initialise as the parsed number chars the change if indexes differ
	var actualFirstNum, actualLastNum = parsedFirstNumber.Number, parsedLastNumber.Number
	// Now to compare indexes
	for _, match := range parsedStringNumbers {
		if match.Index < parsedFirstNumber.Index {
			actualFirstNum = match.Number
		}
		if match.Index > parsedLastNumber.Index {
			actualLastNum = match.Number
		}
	}
	
	fmt.Printf("found values of %d and %d", actualFirstNum, actualLastNum)
	// By this point we should have the true first & last
	var firstRune, lastRune = rune(actualFirstNum), rune(actualLastNum)

	//var combo = string(firstRune) + string(lastRune)
	var combo = string([]rune{firstRune, lastRune})
	fmt.Print("Number is ", combo)	

	var cal, err = strconv.Atoi(combo)
	if (err != nil) {
		println("Parsing Error!", err)
	}
	return cal
}

func parseNumbersFromLine(line string) (Pair, Pair) {
	var firstDigit = Pair{-1, -1}
	var lastDigit Pair
	// for part two we need to parse words as ints.
	// this means we will need to take substrings and check the starting byte.
	for index, char := range line {
		if (unicode.IsDigit(char)){
			if (firstDigit.Index == -1) {
				firstDigit.Number = getNumFromChar(char)
				firstDigit.Index = index 
			}
		 	lastDigit.Number = getNumFromChar(char)
			lastDigit.Index = index
		}
	}

	return firstDigit, lastDigit
}

func getNumFromChar(char rune) int {
	var number, err = strconv.Atoi(string(char))			
	if err != nil {
		log.Fatal("Couldnt properly parse a number from a rune")
	}
	return number
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

// for later remember a line could contain to numsAsWords so you need to remove as you find
func parseLineToInts(line string) []Pair {
	var matches []Pair
	
	for numAsWord, num := range numMap {
		// Go until you find no hits
		for {
		// if word contains numAsWord, return num & index
			if strings.Contains(line, numAsWord){
				var index = strings.Index(line, numAsWord)
				// Replacing only one instance as there is a high probability a
				// word could repeat in the same line
				strings.Replace(line, numAsWord, "", 1)

				var match = Pair{index, num }
				matches = append(matches, match) 
			} else {
				break
			}
		}

	}
	return matches
}

type Pair struct {
	Index int
	Number int
}
