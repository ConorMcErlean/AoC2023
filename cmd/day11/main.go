package main

import (
	"adventOfCode23/cmd/common"
	"fmt"
)

func main(){
	file := common.ReadFileFromArgs()
	galaxies, Universe := fileToUniverse(file)
	galaxies2 := fileToBigUniverse(file)
	
	Universe = append(Universe, "")

	pairs := makeGalaxyPairs(galaxies)
	pairs2 := makeGalaxyPairs(galaxies2)
	total, total2 := int64(0), int64(0)

	for _, pair := range pairs {
		pair =	pair.calculateDistance()
		total += pair.distance
	}

	for _, pair := range pairs2 {
		pair =	pair.calculateDistance()
		pair.printPair()
		total2 += pair.distance
	}

	fmt.Printf("\nTotal Distance %v vs %v\n", total, total2)
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
				galaxies = append(galaxies, Galaxy{X: int64(x + gravityModifier), Y: int64(y), Number: galaxyNum})
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

func fileToBigUniverse(file []string) (galaxies []Galaxy) {
	galaxyNum := 1
	gravityModifierX := int64(0)
	gravityScaleFactor := int64(1_000_000)

//	gravityScaleFactor := 1000000 
	var indexesForYGravity []int
	for y := range file[0] {
		if checkColumnEmpty(y, file) {
			indexesForYGravity = append(indexesForYGravity, y)
		}
	}

	
	for x, line := range file {
		yIndex := 0
		gravityModifierY := int64(0)
		galaxyFound := false
		for y, char := range line {
			// Whem condition not met, we have done all the gravity lines
			if yIndex < len(indexesForYGravity){
				if y == indexesForYGravity[yIndex] {
					yIndex++
					gravityModifierY++
				}
			}
				if char != '.' {
				
			//	galaxies = append(galaxies, Galaxy{X: int64(x)-1 + int64(gravityScaleFactor) , Y: int64(y)-1 + int64(gravityScaleFactor) , Number: galaxyNum})
				galaxies = append(galaxies, Galaxy {
					X: ( int64(x)-gravityModifierX ) + ( gravityModifierX * gravityScaleFactor ) , 
					Y: ( int64(y)-gravityModifierY ) + ( gravityModifierY * gravityScaleFactor ) , Number: galaxyNum})
				galaxyNum++
				galaxyFound = true
			}
		}
		if !galaxyFound {
			gravityModifierX++
		}
	}
	return galaxies
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
	distance := int64(0)
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
	fmt.Printf("\n Galaxy %v (%v, %v) to Galaxy %v (%v, %v) distance %v", pair.a.Number,pair.a.X, pair.a.Y, pair.b.Number, pair.b.X, pair.b.Y, pair.distance)
}

type Galaxy struct {
	Number int
	X int64
	Y int64
}

type GalaxyPair struct {
	distance int64
	a Galaxy
	b Galaxy
}
