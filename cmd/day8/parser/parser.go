package parser

import (
	"fmt"
	"strings"
)

func GetDirections(line string) []bool {
	// left = true
	line = strings.TrimSpace(line)
	directions := make([]bool, len(line))

	for index, character := range line {

		switch(character) {
		case 'L':
			directions[index] = true
		case 'R' :
			directions[index] = false
		default :
			fmt.Println("Unusual character:", string(rune(character)))
		}
	}

	return directions
}

func GetMap(lines []string) map[string][]string {
	theMap := make(map[string][]string)
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		start, next := getLine(line)
		theMap[start] = next
	}
	return theMap
}

func getLine(line string) (string, []string) {
	parts := strings.Split(line, " = ")
	nextSteps := strings.Split(parts[1], ", ")
	left := strings.TrimSpace(nextSteps[0])
	right := strings.TrimSpace(nextSteps[1])
	left = strings.ReplaceAll(left, "(", "")
	right = strings.ReplaceAll(right, ")", "")
	nextSteps[0] = left
	nextSteps[1] = right
	return strings.TrimSpace(parts[0]), nextSteps

}
