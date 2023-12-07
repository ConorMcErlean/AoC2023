package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/almanac"
	"fmt"
)

func main() {
	file := common.ReadFileFromArgs()
	seeds, almanac := almanac.GetSeedsAndAlmanac(file)
	
	var seedRanges []Range
	for _, seedMap := range seeds {
		seedRanges = append(seedRanges, Range { Low: seedMap.Low, High: seedMap.High })
	}
//	fmt.Println("Ranges going in", seedRanges)
	locations := FindLocation(seedRanges, almanac)
	lowest := int64( 999999999999 )
	for _, location := range locations {
		if location.Low < lowest {
			lowest = location.Low
		}
	}
	//locations := FindSeedLocations(seeds, almanac)//
	fmt.Printf("\nClosest Location  : %v \n", lowest)
}

func FindLocation(input []Range, almanac [][]almanac.Mapping) []Range {
	currentRange := input

//	custom1 := Range { Low: 57, High: 69 }
//	custom2 := Range { Low: 81, High: 94 }
//	test := []Range {custom1, custom2}
//	fmt.Printf("\nRange at start: %v", test)
//	currentRange = ConvertAlmanacLayer(test, almanac[1])
//	fmt.Printf("\nRange at end: %v", currentRange)

	for _, almanacPart := range almanac {
	//	fmt.Printf("\nRange at start: %v", currentRange)

		currentRange = ConvertAlmanacLayer(currentRange, almanacPart)

//		fmt.Printf("\nRange at end: %v", currentRange)
	}
	return currentRange
}

func ConvertAlmanacLayer(input []Range, almanac []almanac.Mapping) (output []Range) {
	rangeOfSeeds := input
	for _, mapping := range almanac {
		// For each mapping go through all ranges, until none match
		for {
			var anyMatched = false
			for index, seedRange := range rangeOfSeeds {
				found, matched, remainder, extraRem := CheckOverlap(mapping, seedRange)
				
				if found {
					anyMatched = true
				}
				if matched.Low != -1 {
					output = append(output, matched)
				}
				if remainder.Low != -1 {
					// replace with remainder
					rangeOfSeeds[index] = remainder
				} else {
					// Mark for cleanup
					rangeOfSeeds[index] = remainder
				}
				if extraRem.Low != -1 {
					rangeOfSeeds = append(rangeOfSeeds, extraRem)
				}

			//	fmt.Println("\nCheck 3")
			//	fmt.Printf("\nCompared range Seeds: %v - %v to %v - %v", seedRange.Low, seedRange.High, mapping.Low, mapping.High )
			//	fmt.Printf(" | Remaining is %v", rangeOfSeeds)
			//	fmt.Printf(" | Mapped is %v", output)
			}

			// cleanup
			rangeOfSeeds = CleanupList(rangeOfSeeds)

			if !anyMatched {
				break
			}

		}
	}
	//fmt.Printf("\nAdding Remainder %v", rangeOfSeeds)
	// At this point any remaining ranges need added to our output
	for _, seedRange := range rangeOfSeeds {
		output = append(output, seedRange)
	}
	//fmt.Println("Output now", output)
	
	return output
}

func CheckOverlap(mapping almanac.Mapping, seedRange Range) (bool, Range, Range, Range) {
	falseMap := Range { Low: -1, High: -1 }
	var remainder Range
	var mapped Range
	extraRemainder := falseMap
	foundMatch := true

	//fmt.Printf("\nComparing range Seeds: %v - %v to %v - %v", seedRange.Low, seedRange.High, mapping.Low, mapping.High )

	switch {
	// No overlap Can contine without checking
	case ( (mapping.Low > seedRange.High) || (mapping.High < seedRange.Low) ) :
		remainder = seedRange
		mapped = falseMap
		foundMatch = false

	// Entire Area is covered by range
	case mapping.Low <= seedRange.Low && mapping.High >= seedRange.High :
		mapped = Range { 
			Low: seedRange.Low - mapping.Translation,
			High : seedRange.High - mapping.Translation,
		}
		remainder = falseMap
	// Left Side Overlap
	case mapping.Low <= seedRange.Low :
		mapped = Range {
			Low : seedRange.Low - mapping.Translation,
			High : mapping.High - mapping.Translation,
		}
		remainder = Range{
			Low: mapping.High +1,
			High: seedRange.High,
		}
	// Right side Overlap
	case mapping.Low > seedRange.Low && mapping.High >= seedRange.High :
		mapped = Range {
			Low: mapping.Low - mapping.Translation,
			High : seedRange.High - mapping.Translation,
		}
		remainder = Range{
			Low: seedRange.Low,
			High : mapping.Low -1,
		}
	// Mapping takes cutout in middle of range
	case mapping.Low > seedRange.Low :
		mapped = Range {
			Low : mapping.Low - mapping.Translation,
			High : mapping.Low - mapping.Translation,
		}
		remainder = Range{
			Low: seedRange.Low,	
			High : mapping.Low -1,
		}
		extraRemainder = Range { 
			Low: mapping.High +1,
			High : seedRange.High,
		}
	}

	//fmt.Printf(" | Remaining is %v - %v", remainder.Low, remainder.High)
	//fmt.Printf(" | Mapped is %v - %v", mapped.Low, mapped.High)
	//fmt.Println("Found match", foundMatch)
	return foundMatch, mapped, remainder, extraRemainder 
}

func CleanupList(list []Range) (cleanList []Range) {
	for _, item := range list {
		if item.Low != -1 {
			cleanList = append(cleanList, item)
		}
	}
	return cleanList
}

func deleteElementAtIndex(index int, collection []Range ) []Range {
	fmt.Printf("\n Requested to delete item at index %v, for length %v, Range: %v \n ", index, len(collection), collection)
	if len(collection) == 1 {
		return collection[:0]
	}
	// last element to this index
	collection[index] = collection[len(collection)-1]
	// delete last element
	collection = collection[:len(collection)-1]
	return collection
}

type Range struct {
	High int64
	Low int64
}

