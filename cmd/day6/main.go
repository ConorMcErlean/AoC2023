package main

import (
	"adventOfCode23/cmd/day6/reader"
	"fmt"
	"math"
)

func main() {
	time, distance := reader.ReadFileLongValues()
	lower, upper := doMath(time, distance)
	if CheckIfMatchesInt(upper) {
		upper -= 1
	}
	fmt.Printf("\n True lower %v, true upper %v", lower, upper)
		
	solutions := int(upper) - int(lower)

	//solutions := CheckPossibleSolutions(time, distance)
	fmt.Printf("\n Solutions in range %v - %v with diff of %v", int(lower +1), int(upper), solutions)
	
	fmt.Printf("\nAnswer: %v\n", solutions)

}

func part1() {
	total := 1
	timesToDistances := reader.ReadFile()
	for time, distance := range timesToDistances {
		lower, upper := doMath(time, distance)
		if CheckIfMatchesInt(upper) {
			upper -= 1
		}
		fmt.Printf("\n True lower %v, true upper %v", lower, upper)
		
		solutions := int(upper) - int(lower)


		//solutions := CheckPossibleSolutions(time, distance)
		fmt.Printf("\n Solutions in range %v - %v with diff of %v", int(lower +1), int(upper), solutions)
		total *= solutions
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

func CheckIfMatchesInt(num float64) bool {
	numAsInt := int(num)
	if (num - float64(numAsInt) ) > 0 { 
		return false 
	}
	return true
}

func doMath(time int, score int) (float64, float64){
	// score +1 = chargeTime^2 + chargeTime * time
	// time = (score + chargeTime ^2 ) / chargeTime
	part1 := (time * time) - (4 * score)
	upperBound := (float64(time) + math.Sqrt(float64(part1)))/2
	lowerBound := (float64(time) - math.Sqrt(float64(part1)))/2
	return lowerBound, upperBound
}
