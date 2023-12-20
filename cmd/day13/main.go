package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
	"slices"
	"strings"
)

func main() {
	file := common.ReadFileFromArgsWithEmptyLines()
	maps := parseToMaps(file)
	total :=0
	for i, mapp := range maps {
		count := 0
				
		fmt.Printf("\n Checking map %v of %v", i+1, len(maps))
		mirrorPointHor, foundH := findHorizontalReflectionWithSmudge(mapp)
		mirrorPointVert, foundV := findVerticalReflectionWithSmudge(mapp)
		//fmt.Printf("\nSmudged Reflections: Horizontal (%v) %v, Vertical (%v) %v", foundH, mirrorPointHor, foundV, mirrorPointVert)
		fmt.Print("\nSmudged Reflections:")
		if foundH {
			fmt.Printf(" Horizontal: %v", mirrorPointHor)
		}
		if foundV {
			fmt.Printf(" Vertical : %v", mirrorPointVert)
		}

		
		if foundH && foundV {
			if mirrorPointHor < mirrorPointVert {
				mirrorPointVert = 0
			} else {
				mirrorPointHor = 0 
			}
		}
		

		if foundH {
			count = mirrorPointHor
		}

		if foundV {
			count += (100* mirrorPointVert)
		}
		total += count
	}
	fmt.Printf("\n\nTotal %v\n", total)
}

func parseToMaps(input []string) [][]string {
	var maps [][]string 
	var currentMap []string

	for i, line := range input {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			currentMap = append(currentMap, line)
		} else {
			maps = append(maps, currentMap)
			currentMap = make([]string, 0)

		}
		if i == len(input)-1 && len(currentMap) != 0 {
			maps = append(maps, currentMap)
		}
	}
	return maps
}


func findHorizontalReflectionWithSmudge(input []string) (int, bool) {
	var horizontal, lastHorizontal []rune
	possibleReflections := make(map[int]bool)
	var reflections []int 

	for i := range input[0] {
		horizontal = createColumn(i, input)	
		if i == 0 {
			lastHorizontal = horizontal
			continue 
		}
		equal, neededSmudge := CheckListOfCharacters(horizontal, lastHorizontal)
		if equal {
			possibleReflections[i] = neededSmudge
		}

		lastHorizontal = horizontal
	}

	for reflection, smudgeUsed := range possibleReflections {
		left, right := reflection -2, reflection+1
		for {
			if !(left >= 0) || !(right < len(input[0])) {
				// Must be exactly one smudge
				if smudgeUsed {
					reflections = append(reflections, reflection)
				}
				break
			}


			leftColumn, rightColumn := createColumn(left, input), createColumn(right, input)
		//	fmt.Printf("\nComparing\n%v\n%v", leftColumn, rightColumn)
			equal, neededSmudge := CheckListOfCharacters(leftColumn, rightColumn)
	//		if reflection == 5 {
	//			fmt.Printf("\n equal %v, needed smudge %v", equal, neededSmudge)
	//		}
			if(!smudgeUsed) {
				smudgeUsed = neededSmudge
			} else if smudgeUsed && neededSmudge {
				break
			}	
			if !equal {
				break
			}

			left--
			right++

		}
	}
	fmt.Println("Horizontal")
	
	return ReturnCorrect(reflections, input)
}

func findVerticalReflectionWithSmudge(input []string) (int, bool) {
	var lastline string
	possibleReflections := make(map[int]bool)
	var reflections []int 

	for index, line := range input {
		if index == 0 {
			lastline = line
			continue
		}

		equal, neededSmudge := CheckListOfCharacters([]rune(lastline), []rune(line))

		if equal {
			possibleReflections[index] = neededSmudge
		}
		lastline = line
	}

	for reflection, usedSmudge := range possibleReflections {
		// Aleady Compared -1 & 0
		up, down := reflection-2, reflection+1
		for {
			if !(up >= 0) || !(down < len(input) ) {
				if usedSmudge {
					reflections = append(reflections, reflection)
				}
				break
			}

			upString, downString := []rune(input[up]), []rune(input[down])
		//	fmt.Printf("\nComparing\n%v\n%v", upString, downString)
			equal, neededSmudge := CheckListOfCharacters(upString, downString)
			if(!usedSmudge) {
				usedSmudge = neededSmudge
			} else if usedSmudge && neededSmudge {
				break
			}	
			if !equal {
				break
			}

			down++
			up --
		}
	}
	fmt.Println("Vertical")
	
	return ReturnCorrect(reflections, input)

}

func ReturnCorrect(reflections []int, input []string) (int, bool) {
	if len(reflections) == 0 {
		return -1, false
	} 
	if len(reflections) > 1 {
		print("\nSomething went wrong, Printing challenge\n--")
		for i := range input[0] {
			num := i +1
			if num > 9 {
				num -= 10
			}
			fmt.Print(num)

		}
		fmt.Print("\n")
		for i, line := range input {
			num := i +1
			if num > 9 {
				num -= 10
			}
			fmt.Printf("%v ",num)

			fmt.Println(line)
		}

		fmt.Printf("\n\nReflection point %v\n\n", reflections)
		panic("Should only be one!")
	}

	return reflections[0], true

}

func createColumn(index int, input []string) []rune {
	column := make([]rune, len(input))

	for row := range input { 
		column[row] = rune(input[row][index])
	}	

	return column
}

func CheckListOfCharacters(in1 []rune, in2 []rune) (equal bool, smudge bool) {
	// check for equality or exactly one difference
	equal, smudge = true, false
	for i := 0; i < len(in1); i++ {
		if in1[i] == in2[i] {
			continue
		} else if !smudge {
			smudge = true
		} else {
			return false, false
		}
	}
	return equal, smudge
}

func CheckListOfCharactersDebug(in1 []rune, in2 []rune) (equal bool, smudge bool) {
	// check for equality or exactly one difference
	fmt.Printf("\nComparing\n%v\n%v", in1, in2)
	equal, smudge = true, false
	for i := 0; i < len(in1); i++ {
		if in1[i] == in2[i] {
			continue
		} else if !smudge {
			fmt.Printf("\n Smudge used on index %v", i )
			smudge = true
		} else {
			return false, false
		}
	}
	fmt.Println("Return as a match")
	return equal, smudge
}

// Part 1

func findHorizontalReflection(input []string) (int, bool) {
	var horizontal, lastHorizontal []rune
	var possibleReflections []int 

	for i := range input[0] {
		horizontal = createColumn(i, input)	
		if i == 0 {
			lastHorizontal = horizontal
			continue 
		}
		if slices.Equal(horizontal, lastHorizontal) {
			possibleReflections = append(possibleReflections, i)
		}
		lastHorizontal = horizontal
	}
	if len(possibleReflections) == 0 {
		fmt.Println("No possible horizontal")
	}
	for _, reflection := range possibleReflections {
		left, right := reflection -2, reflection+1
		for {
			if !(left >= 0) || !(right < len(input[0])) {
				return reflection, true
			}


			leftColumn, rightColumn := createColumn(left, input), createColumn(right, input)
		//	fmt.Printf("\nComparing\n%v\n%v", leftColumn, rightColumn)
			if !slices.Equal(leftColumn, rightColumn ) {
				break
			}

			left--
			right++

				}
	}
	return -1, false
}


func findVerticalReflection(input []string) (int, bool) {
	var lastline string
	var possibleReflections []int 

	for index, line := range input {
		if index == 0 {
			lastline = line
			continue
		}
		if strings.EqualFold(lastline, line) {
			possibleReflections = append(possibleReflections, index)
		}
		lastline = line
	}

	for _, reflection := range possibleReflections {
		up, down := reflection -1, reflection
		for {
			if !(up >= 0) || !(down < len(input) ) {
				return reflection, true
			}

			upString, downString := input[up], input[down]
		//	fmt.Printf("\nComparing\n%v\n%v", upString, downString)
			if !strings.EqualFold(upString, downString) {
				break
			}
			down++
			up --
		}
	}
	return -1, false
}
