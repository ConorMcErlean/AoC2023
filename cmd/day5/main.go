package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/almanac"
	"fmt"
	"slices"
)

func main() {
	file := common.ReadFileFromArgs()
	seeds, almanac := almanac.GetSeedsAndAlmanac(file)
	var locations []int64
	for _, seed := range seeds {
		soil := ConvertToDestination(seed, almanac["seed-to-soil"])
		fertilizer := ConvertToDestination(soil, almanac["soil-to-fertilizer"])
		water := ConvertToDestination(fertilizer, almanac["fertilizer-to-water"])
		light := ConvertToDestination(water, almanac["water-to-light"])
		temperature := ConvertToDestination(light, almanac["light-to-temperature"])
		humidity := ConvertToDestination(temperature, almanac["temperature-to-humidity"])
		location := ConvertToDestination(humidity, almanac["humidity-to-location"])
		fmt.Printf("\nSeed %v Goes in location %v\n", seed, location)
		locations = append(locations, location)
	}

	fmt.Printf("Closest Location  : %v ", slices.Min(locations))
}

func ConvertToDestination(input int64, almanacPart map[int64]int64) int64 {
	destination, exists := almanacPart[input]
	if !exists {
		destination = input
	}
	return destination
}
