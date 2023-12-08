package filereading

import (
	"slices"
	"strings"
)

// Parsing File to Chunks
func BreakInputIntoComponents(file []string ) [][]string  {
	var linesInMap []string
	lastIndex := len(file) -1
	var individualMaps [][]string
	

	for index, line := range file {
		if strings.Contains(line, "map") {
			if len(linesInMap) > 0 {
				// assign
				individualMaps = writeString(individualMaps, linesInMap)
			}
			// New Map
			// empty the slice
			linesInMap = linesInMap[:0]
		} else if strings.Contains(line, "seeds:") {
			lines := strings.Split(line, ":")
			linesInMap = append(linesInMap, lines[1])
			individualMaps = writeString(individualMaps, linesInMap)
			// empty lines
			linesInMap = linesInMap[:0] 
		} else {
			linesInMap = append(linesInMap, line)
		}	
		if (index == lastIndex) {
			individualMaps = writeString(individualMaps, linesInMap)
		}
	}

	return individualMaps
}

func ReadHeader(line string) string {
	header := strings.Replace(line, "map:", "", 1)
	header = strings.TrimSpace(header)
	return header
}

func writeString(mapOfIndexes [][]string, lines []string) [][]string {
	// Spent forever before realising I was previously assingning the Ref to the map,
	// but the underlying values were changing!
	linesForKey := slices.Clone(lines)
	mapOfIndexes = append(mapOfIndexes, linesForKey )
	return mapOfIndexes
}
