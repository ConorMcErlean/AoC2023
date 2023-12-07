package main

import (
	"adventOfCode23/cmd/day6/reader"
	"fmt"
)

func main() {
	total := 1
	timesToDistances := reader.ReadFile()
	for time, distance := range timesToDistances {
		solutions := CheckPossibleSolutions(time, distance)
		fmt.Printf("\n Solutions include %v", solutions)
		total *= len(solutions)
	}
	fmt.Printf("\nAnswer: %v\n", total)

}

func CheckPossibleSolutions(time int, distance int) (solutions []int) {
	// distance = (chargetime) x (time-chargetime)
	for chargeTime := 1; chargeTime <=time; chargeTime++ {
		if ( chargeTime * (time - chargeTime) ) > distance {
			solutions = append(solutions, chargeTime)
		}
	}
	return solutions
}
