package main

import (
	"adventOfCode23/cmd/common"
	"adventOfCode23/cmd/day5/almanac"
)

func main() {
	file := common.ReadFileFromArgs()
	almanac.FindMap(file)
}
