package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strings"
)

func main() {
	values := GetValues()
	total := 0
	for _, valuelist := range values {
		next := FindNext(valuelist)
		fmt.Printf("\nNext value in list %v : %v\n", valuelist, next)
		total += next
	}
	fmt.Printf("\n Total of Next Values : %v \n\n", total)	
}

func GetValues() [][]int {
	lines := common.ReadFileFromArgs()
	values := make([][]int, len(lines))
	for index, line := range lines {
		values[index] = lineToValues(line)
	}
	return values
}

func FindNext(values []int) int {
	var differenceLayers [][]int
	foundEnd := false
	var differences []int
	differenceLayers = append(differenceLayers, values)
	for {
		differences, foundEnd = getDifferences(values)
		differenceLayers = append(differenceLayers, differences)
		values = differences

		if foundEnd {
			break
		}
	}

	// Now find next number
	total := 0
	fmt.Printf("Difference Layers %v", differenceLayers)
	for i := len(differenceLayers)-1; i >= 0; i-- {
		layer := differenceLayers[i]
		total += layer[len(layer)-1]
	}

	return total

}

func getDifferences(values []int) ([]int, bool) {
	differences := make([]int, len(values) -1)
	var current, next int

	for i := 0; i < len(values) -1; i++ {
		current = values[i]
		next = values[i +1]
		differences[i] = next - current
	}

	for _, difference := range differences {
		if difference != 0 {
			return differences, false
		}
	}
	return differences, true
}

func lineToValues(line string) (values []int) {
	sections := strings.Split(line, " ")
	for _, section := range sections {
		if section == " " {
			continue 
		}
		values = append(values, common.StringToInt(section))
	}
	return values
}
