package reader

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strings"

)

func ReadFile() map [int]int {
	times, distances := ReadAndBreakUpIntoComponents() 

	timeToDistance := make(map[int]int)

	for index, timeValue := range times {
		timeToDistance[common.StringToInt(timeValue)] = common.StringToInt(distances[index])
	}
	
	fmt.Printf("times to distances %v", timeToDistance)
	return timeToDistance
}

func ReadFileLongValues() (int, int) {
	times, distances := ReadAndBreakUpIntoComponents() 
	timeString := ""
	distanceString :=""
	for index, timeValue := range times {
		timeString = timeString + timeValue
		distanceString = distanceString + distances[index]
	}
	
	time := common.StringToInt(timeString)
	distance := common.StringToInt(distanceString)
	return time, distance
}

func ReadAndBreakUpIntoComponents() ( []string, []string ) {
	file := common.ReadFileFromArgs()
	time := strings.Split(file[0], ":")[1]
	distance := strings.Split(file[1], ":")[1]
	times := strings.Split(time, "  ")
	distances := strings.Split(distance, "  ")
	distances = RemoveSpaces(distances)
	times = RemoveSpaces(times)
	return times, distances
}


func RemoveSpaces(in []string) (out []string) {
	for _, item := range in {
		item = strings.TrimSpace(item)
		if item != "" {
			out = append(out, item)
		}
	}
	return out
}
