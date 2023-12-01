package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main(){
	var arguements = os.Args
	if len(arguements) < 2 {
		log.Fatal("We really need a file to do anything")
	}
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
	var firstDigit, lastDigit string
	for _, char := range input {
		if (unicode.IsDigit(char)){
			if (firstDigit == "") {
				firstDigit = string(char)
			}
		 	lastDigit = string(char)
		}
	}
	var combo = firstDigit + lastDigit
	var cal, err = strconv.Atoi(combo)
	if (err != nil) {
		println("Parsing Error!", err)
	}
	return cal
}
