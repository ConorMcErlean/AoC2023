package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strings"
)

 func main() {
	 // I needed help here as my brain is fried
	file := common.ReadFileFromArgs()
	total :=0
	for _, line := range file {
		parts := strings.Split(line, " ")
		springs := parts[0]
		counts := getCounts(parts[1]) 
		possibilities := count(springs, counts)
		total += possibilities
		fmt.Printf("\nPossibilities: %v", possibilities)
	}

	fmt.Printf("\nPossible options: %v\n", total)



}

func getCounts(input string) []int {
	var counts []int
	parts := strings.Split(input, ",")
	for _, part := range parts {
		counts = append(counts, common.StringToInt(part))
	}
	return counts
}

func count(springs string, counts []int) int  {
	result := 0
	if springs == "" {
		if len(counts) == 0 {
			return 1
		}
		return 0
	}
	if len(counts) == 0 {
		if strings.Contains(springs, "#") {
			return 0 
		} 
		return 1
	}

	if springs[0] == '.' || springs[0] == '?' {
		result += count(springs[1:], counts)
	}
	// 5:47 Differs
	if springs[0] == '#' || springs[0] =='?' {
		if counts[0] <= len(springs) && !strings.Contains(springs[:counts[0]], ".") && ( counts[0] == len(springs) || springs[counts[0]] != '#' ){
			next := counts[0] +1
			if next > len(springs) {
				result += count("", counts[1:])
			} else {
				result += count(springs[next:], counts[1:])
			}
		} 
	}
	return result
}
