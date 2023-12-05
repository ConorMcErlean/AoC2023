package almanac

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/filereading"
	"fmt"
	"strings"
)

func FindMap(file []string){
	var seeds []int
	almanac := make(map[string]map[int]int)
	almanacParts := filereading.BreakInputIntoComponents(file)


	for key, value := range almanacParts {
		fmt.Printf("\nMap %v will contain %v\n", key, value)
		if key == "seeds" {
			seeds = getSeeds(value)
		} else {
			almanac[key] = ConvertLinesToMap(value)
		}
	}
	fmt.Println("Seeds:", seeds)

}

func ConvertLinesToMap(input []string) map[int]int {
	almanacMap := make(map[int]int)

	for index, line := range input {
		if (index == 0) {
			continue 
		}
		lineVals := strings.Split(line, " ")
		destination := getValueFrom(lineVals, 0)
		source := getValueFrom(lineVals, 1)
		length := getValueFrom(lineVals, 2)
		for i := 0; i < length; i++ {
			almanacMap[source + i] = destination + i 
		}
	}
	fmt.Println("\nMap has values:", almanacMap)
	return almanacMap
}

func getSeeds(lines []string)(seeds []int) {
	for _, line := range lines {
		fmt.Println("SL:", line)
		if len(line) == 0 {
			continue
		}
		sections := strings.Split(line, ":")
		valuesString := sections[1]
		fmt.Println("Seed Values:", valuesString)
		for _, value := range strings.Split(valuesString, " ") {
			fmt.Printf("Seeds array parsing %v", value)
			seeds = append(seeds, common.StringToInt( strings.TrimSpace(value)))
		}
	}
	return seeds
}

func getValueFrom(values []string, index int) int {
	value := strings.TrimSpace(values[index]) 
//	fmt.Printf("\n Going to parse to Value: %v", value)
	number := common.StringToInt( value )
	return number
}


