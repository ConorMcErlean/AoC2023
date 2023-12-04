package scratchcardreader

import (
	"strings"
	"strconv"
)

func ReadScratchCard(line string) (winners []int, cardNumbers []int) {
	card := strings.Split(line, ":")[1]
	components := strings.Split(card, "|")
	winSide := strings.TrimSpace(components[0])
	gameSide := strings.TrimSpace(components[1])
	winners = readValues(winSide)
	cardNumbers = readValues(gameSide)

	return winners, cardNumbers
}

func readValues(values string) (numbers []int) {
	for _, value := range strings.Split(values, " ") {
		value = strings.TrimSpace(value)
		val, err := strconv.Atoi(value)
		
		if (err == nil ) {
			numbers = append(numbers, val)
		}
	}
	return numbers
}
