package almanac

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/filereading"
	"strings"
)

func GetSeedsAndAlmanac(file []string) (seeds []Mapping, almanac map[string][]Mapping) {
	almanac = make(map[string][]Mapping)
	almanacParts := filereading.BreakInputIntoComponents(file)
	
	for key, value := range almanacParts {
		if key == "seeds" {
			seeds = getSeeds(value)
		} else {
			almanac[key] = ConvertLinesToMap(value)
		}
	}
	return seeds, almanac
}

func ConvertLinesToMap(input []string) (mappings []Mapping) {
	for _, line := range input {
		lineVals := strings.Split(line, " ")
		destination := getValueFrom(lineVals, 0)
		source := getValueFrom(lineVals, 1)
		length := getValueFrom(lineVals, 2)
		// Previous version of loop was insane
		mapping := Mapping { Low: source, High: source + length, Translation: source - destination }
		mappings = append(mappings, mapping)
	}
	return mappings
}

func getSeeds(lines []string) []Mapping {
	var seeds []int64
	var seedMappings []Mapping
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
	for i := 0; i < len(seeds); i += 2 {
		length := seeds[i+1]
		start := seeds[i]
		// -1 Because length includes first number
		mapping := Mapping { Low: start, High: start + length -1 } 	
		seedMappings = append(seedMappings, mapping)
	}
	return seedMappings
}

func getValueFrom(values []string, index int) int64 {
	value := strings.TrimSpace(values[index]) 
	number := int64(common.StringToInt( value ))
	return number
}

type Mapping struct {
	Low int64
	High int64
	Translation int64
}
