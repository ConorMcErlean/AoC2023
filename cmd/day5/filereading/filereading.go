package filereading

import (
	"slices"
	"strings"
)

// Parsing File to Chunks
func BreakInputIntoComponents(file []string ) map[string][]string  {
	var linesInMap []string
	lastIndex := len(file) -1
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
			writeMap(individualMaps, lastHeader, linesInMap)
		}
	}

	return individualMaps
}

func ReadHeader(line string) string {
	header := strings.Replace(line, "map:", "", 1)
	header = strings.TrimSpace(header)
	return header
}

func writeMap(mapOfIndexes map[string][]string, key string, lines []string) {
	// Spent forever before realising I was previously assingning the Ref to the map,
	// but the underlying values were changing!
	linesForKey := slices.Clone(lines)
	mapOfIndexes[key] = linesForKey
}
