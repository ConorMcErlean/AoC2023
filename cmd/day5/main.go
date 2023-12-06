package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/almanac"
	"fmt"
	"slices"
	"sync"
)

func main() {
	file := common.ReadFileFromArgs()
	seeds, almanac := almanac.GetSeedsAndAlmanac(file)
	locations := FindSeedLocations(seeds, almanac)
	fmt.Printf("\nClosest Location  : %v \n", slices.Min(locations))
}

func ConvertToDestination(input int64, almanacPart []almanac.Mapping) (destination int64) {
	var exists = false
	for _, mapping := range almanacPart {
		if (input >= mapping.Low) && (input <= mapping.High) {
			destination = input - mapping.Translation
			exists = true
			break
		}
	}

	if !exists {
		destination = input
	}
	return destination
}

func FindSeedLocations(seeds []almanac.Mapping, almanac map[string][]almanac.Mapping) (locations []int64) {
	for _, seedMapping := range seeds {
		location := FindLowestLocation(seedMapping, almanac)	
		locations = append(locations, location)
	}
	fmt.Println(locations)
	return locations
}

func FindLowestLocation(seed almanac.Mapping, almanac map[string][]almanac.Mapping) (location int64 ) {
	// Because the mappings along the way could result in any value,
	// need to check all seeds
	locations := SafeSlice { v: make([]int64, seed.High - seed.Low +1 ) }
	var locationsForRange []int64
	var wg sync.WaitGroup
	index := 0
	for i := seed.Low; i <= seed.High; i++ {
		wg.Add(1)
		go func(seedNum int64, sliceIndex int) {
			locations.Add( SeedToLocation(seedNum, almanac), sliceIndex )
			wg.Done()
		}(i, index)
		index++
	}
	wg.Wait()
	locationsForRange = locations.Get()
	lowest := slices.Min(locationsForRange)
	fmt.Printf("\nFor range of seeds %v - %v, the closest location was %v", seed.Low, seed.High, lowest)
	return lowest
}

func SeedToLocation(seed int64, almanac map[string][]almanac.Mapping) int64 {
	soil := ConvertToDestination(seed, almanac["seed-to-soil"])
	fertilizer := ConvertToDestination(soil, almanac["soil-to-fertilizer"])
	water := ConvertToDestination(fertilizer, almanac["fertilizer-to-water"])
	light := ConvertToDestination(water, almanac["water-to-light"])
	temperature := ConvertToDestination(light, almanac["light-to-temperature"])
	humidity := ConvertToDestination(temperature, almanac["temperature-to-humidity"])
	location := ConvertToDestination(humidity, almanac["humidity-to-location"])
//	fmt.Printf("\nSeed %v relates to destination %v\n", seed, location)
	return location
}

type SafeSlice struct {
	mu sync.Mutex
	v []int64
}

func (c *SafeSlice) Add(value int64, index int) {
	c.mu.Lock()
	c.v[index] = value
	c.mu.Unlock()
}

func (c *SafeSlice) Get() []int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
