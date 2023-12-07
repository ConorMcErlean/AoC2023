package reader

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"strings"

)

func ReadFile() map [int]int {
	file := common.ReadFileFromArgs()
	time := strings.Split(file[0], ":")[1]
	distance := strings.Split(file[1], ":")[1]
	timeToDistance := make(map[int]int)
	times := strings.Split(time, "  ")
	distances := strings.Split(distance, "  ")
	distances = RemoveSpaces(distances)
	times = RemoveSpaces(times)
	for index, timeValue := range times {
		timeToDistance[common.StringToInt(timeValue)] = common.StringToInt(distances[index])
	}
	
	fmt.Printf("times to distances %v", timeToDistance)
	return timeToDistance
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
