package almanac

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/filereading"
	"fmt"
	"strings"
)

func GetSeedsAndAlmanac(file []string) (seeds []int64, almanac map[string]map[int64]int64) {
	almanac = make(map[string]map[int64]int64)
	almanacParts := filereading.BreakInputIntoComponents(file)


	for key, value := range almanacParts {
		//fmt.Printf("\nMap %v will contain %v\n", key, value)
		if key == "seeds" {
			seeds = getSeeds(value)
		} else {
			almanac[key] = ConvertLinesToMap(value)
		}
	}
	fmt.Println("Seeds:", seeds)
	return seeds, almanac
}

func ConvertLinesToMap(input []string) map[int64]int64 {
	almanacMap := make(map[int64]int64)

	for index, line := range input {
		if (index == 0) {
			continue 
		}
		lineVals := strings.Split(line, " ")
		destination := getValueFrom(lineVals, 0)
		source := getValueFrom(lineVals, 1)
		length := getValueFrom(lineVals, 2)
		fmt.Println("Got all the values for a line, now to build the map")
		for i := int64(0); i < int64(length); i++ {
			almanacMap[source + i] = destination + i 
		}
		fmt.Println("Built that line!")
	}
	return almanacMap
}

func getSeeds(lines []string)(seeds []int64) {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, value := range strings.Split(line, " ") {
			if (len(value) == 0) || (value == " "){
				// handle any incorrect values from split
				continue
			}
			seeds = append(seeds, int64( common.StringToInt( strings.TrimSpace(value))))
		}
	}
	return seeds
}

func getValueFrom(values []string, index int) int64 {
	value := strings.TrimSpace(values[index]) 
//	fmt.Printf("\n Going to parse to Value: %v", value)
	number := int64(common.StringToInt( value ))
	return number
}


