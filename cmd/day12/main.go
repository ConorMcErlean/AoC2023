package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strconv"
	"strings"
)

var cache map[CacheKey]int64
	 
 func main() {
	 // I needed help here as my brain is fried

	file := common.ReadFileFromArgs()
	total := int64(0)
	for _, line := range file {
		cache = make(map[CacheKey]int64)

		//if i > 5 {	break}
		parts := strings.Split(line, " ")
		springs := parts[0]
		numValues := parts[1]
		for i := 0; i < 4; i++ {
			springs = springs + "?" + parts[0]
			numValues = (numValues + "," + parts[1])

		}
		counts := getCounts(numValues) 
		//fmt.Printf("\norig: %v : %v \nunfolded: %v : %v", parts[0], getCounts(parts[1]), springs, numValues)
		possibilities := count(springs, counts)
		total += possibilities

		fmt.Printf("\nInput: %v arrangements %v", parts[0], possibilities)
		//fmt.Printf("\nPossibilities: %v", possibilities)
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

func count(springs string, counts []int) int64  {
	result := int64(0)

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

	cacheFlattenedCount := ""
	for _, num := range counts {
		cacheFlattenedCount = cacheFlattenedCount + "," + strconv.Itoa(num)
	}
	key := CacheKey{Springs: springs, Nums: cacheFlattenedCount }
	value, exist := cache[key]

	if exist {
		return value
	}

	if springs[0] == '.' || springs[0] == '?' {
		result += count(springs[1:], counts)
	}
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
	cache[key] = result
	return result
}

type CacheKey struct {
	Springs string
	Nums string
}
