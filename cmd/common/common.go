package common

import (
	"log"
	"os"
	"bufio"
)

func ReadFile(filename string) []string {
	var fileLines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fileLines = append(fileLines, scanner.Text())
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatalf("Something went wrong scanning: %v", err)
	}
	return fileLines
}

func ReadFileFromArgs() []string {
	var arguements = os.Args
	if len(arguements) < 2 {
		log.Fatal("We really need a file to do anything")
	}
	return ReadFile(arguements[1])
}
