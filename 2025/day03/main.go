package main

import (
	"adventofcode/util"
	"adventofcode/util/grid"
	"fmt"
)

func main() {
	//data := getData("sample_input.txt")
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part1(grid grid.Grid) int {
	result := 0

	for i := range grid.Height {
		result += maxJoltagePart1(grid, i*grid.Width, grid.Width)
	}

	return result
}
func maxJoltagePart1(grid grid.Grid, start, length int) int {
	maxLocation1 := firstMaxValueLocation(grid.Grid, start, length-1)
	maxLocation2 := firstMaxValueLocation(grid.Grid, maxLocation1+1, length-(maxLocation1-start)-1)

	return grid.Grid[maxLocation1]*10 + grid.Grid[maxLocation2]
}

func part2(grid grid.Grid) int {
	result := 0

	for i := range grid.Height {
		result += maxJoltagePart2(grid, i*grid.Width, grid.Width)
	}

	return result
}
func maxJoltagePart2(grid grid.Grid, start, length int) int {
	result := 0
	startLocation := start

	for i := range 12 {
		maxLocation := firstMaxValueLocation(grid.Grid, startLocation, length-(startLocation-start)-(12-i-1))

		multiplier := util.IntPow(10, 11-i)
		result += grid.Grid[maxLocation] * multiplier
		startLocation = maxLocation + 1
	}

	return result
}

func firstMaxValueLocation(values []int, start, length int) int {
	result := start

	for i := range length {
		value := values[start+i]
		if value > values[result] {
			result = start + i
		}
	}

	return result
}

func getData(filename string) grid.Grid {
	lines := util.GetFileLines(filename)

	mapping := make(map[rune]int)
	grid.GetNumericMapping(mapping)

	return grid.NewFromData(lines, mapping)
}
