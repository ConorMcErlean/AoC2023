package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
)

func main(){
	file := common.ReadFileFromArgs()
	galaxies, Universe := fileToUniverse(file)
	pairs := makeGalaxyPairs(galaxies)
	total := 0

	for _, line := range Universe {
		fmt.Println(line)
	}
	for _, galaxy := range galaxies {
		fmt.Println(galaxy)
	}
	for _, pair := range pairs {
		pair =	pair.calculateDistance()
		pair.printPair()
		total += pair.distance
	}

	fmt.Printf("\nTotal Distance %v\n", total)
}

func fileToUniverse(file []string) (galaxies []Galaxy, Universe []string) {
	galaxyNum := 1
	gravityModifier := 0

	newFile := make([]string, len(file))

	for y := range file[0] {
		for x := range file {
			if checkColumnEmpty(y, file) {
				newFile[x] = newFile[x] + string( rune(file[x][y]) )
				newFile[x] = newFile[x] + string( rune(file[x][y]) )
			} else {
				newFile[x] = newFile[x] + string( rune(file[x][y]) )
			}
		}
	}

	file = newFile

	for x, line := range file {
		galaxyFound := false
		for y, char := range line {
			if char != '.' {
				galaxies = append(galaxies, Galaxy{X: x + gravityModifier, Y: y, Number: galaxyNum})
				galaxyNum++
				galaxyFound = true
			}
		}
		if galaxyFound {
			Universe = append(Universe, line)
		} else {
			gravityModifier++
			Universe = append(Universe, line)
			Universe = append(Universe, line)
		}
	}
	return galaxies, Universe
}

func checkColumnEmpty(y int, file []string) bool {
	for x:= 0; x < len(file); x++ {
		if file[x][y] != '.' {
			return false 
		}
	}
	return true
}

func makeGalaxyPairs(galaxies []Galaxy) []GalaxyPair {
	var pairs []GalaxyPair
	for i := 0; i < len(galaxies); i++ {
		for j := i+1; j < len(galaxies); j++ {
			pairs = append(pairs, GalaxyPair{ a: galaxies[i], b: galaxies[j] })
		}
	}
	return pairs
}

func (pair GalaxyPair) calculateDistance() GalaxyPair {
	// Make X's match
	distance := 0
	//pair.a.X
	if pair.a.X > pair.b.X {
		distance += (pair.a.X - pair.b.X)	
	} else {
		distance += (pair.b.X - pair.a.X)
	}

	if pair.a.Y > pair.b.Y {
		distance += (pair.a.Y - pair.b.Y)	
	} else {
		distance += (pair.b.Y - pair.a.Y)
	}
	pair.distance = distance
	return pair
}

func (pair GalaxyPair) printPair() {
	fmt.Printf("\n Galaxy %v to Galaxy %v distance %v", pair.a.Number, pair.b.Number, pair.distance)
}

type Galaxy struct {
	Number int
	X int
	Y int
}

type GalaxyPair struct {
	distance int
	a Galaxy
	b Galaxy
}
