package filereading

import (
	"fmt"
	"strings"
)

// Parsing File to Chunks 
func BreakInputIntoComponents(file []string ) map[string][]string  {
	var linesInMap []string
	lastIndex := len(file) -1
	//individualMaps := InitialiseStringMap()
	individualMaps := make(map[string][]string)

	var lastHeader string

	for index, line := range file {
		if strings.Contains(line, "map") {
			if len(linesInMap) > 0 {
				// assign
				writeMap(individualMaps, lastHeader, linesInMap)
			}
			// New Map
			lastHeader = ReadHeader(line)
			// empty the slice
			linesInMap = linesInMap[:0]
		} else if strings.Contains(line, "seeds:") {
			fmt.Println("Apparently this line matches with seeds", line)
			lastHeader = "seeds"
			lines := strings.Split(line, ":")
			linesInMap = append(linesInMap, lines[1])
			writeMap(individualMaps, lastHeader, linesInMap)
			// empty lines
			linesInMap = linesInMap[:0] 
		} else {
			linesInMap = append(linesInMap, line)
		}	
		if (index == lastIndex) {
			fmt.Println("Final write!!!!!!!!!!")
			writeMap(individualMaps, lastHeader, linesInMap)
		}
	}

	fmt.Println("\n -- Sanity Check --")
	for key, value := range individualMaps{
		fmt.Printf("\nMap %v contains %v", key, value)
	}
	return individualMaps
}

func ReadHeader(line string) string {
	header := strings.Replace(line, "map:", "", 1)
	header = strings.TrimSpace(header)
	return header
}

func writeMap(mapOfIndexes map[string][]string, key string, lines []string) {
	fmt.Printf("\nAssigning to map %v, the list %v", key, lines)
	mapOfIndexes[key] = lines

}

//func InitialiseStringMap() map[string][]string {
//	mapOfIndexes := make(map[string][]string)
//	mapOfIndexes["seeds"] = nil
//	mapOfIndexes["seed-to-soil"] = nil
//	mapOfIndexes["soil-to-fertilizer"] = nil
//	mapOfIndexes["fertilizer-to-water"] = nil
//	mapOfIndexes["water-to-light"] = nil
//	mapOfIndexes["light-to-temperature"] = nil
//	mapOfIndexes["temperature-to-humidity"] = nil
//	mapOfIndexes["humidity-to-location"] = nil
//	return mapOfIndexes	
//}
