package main

import (
	"adventofcode/util"
	"adventofcode/util/grid"
	"fmt"
	"log"
	"runtime/debug"
)

func RuntimeAssert(condition bool) {
	if !condition {
		debug.PrintStack()
		log.Fatalln("Assertion failed")
	}
}

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(&data))
	fmt.Printf("Part 2: %d\n", part2(&data))
}
func part1(grid *grid.Grid) int {
	result := 0

	for cell, value := range grid.Grid {
		if value != 0 {
			continue
		}

		target := make(map[int]bool)
		getTrailHeads(grid, cell, target)
		result += len(target)
	}

	return result
}
func part2(grid *grid.Grid) int {
	result := 0

	for cell, value := range grid.Grid {
		if value != 0 {
			continue
		}

		target := make(map[int]bool)
		result += getTrailHeads(grid, cell, target)
	}

	return result
}
func getTrailHeads(grid *grid.Grid, start int, target map[int]bool) int {
	currentValue := grid.Grid[start]
	if currentValue == 9 {
		target[start] = true
		return 1
	}

	result := 0

	neighbours := grid.GetNeighbours(start)
	for _, neighbour := range neighbours {
		if currentValue+1 == grid.Grid[neighbour] {
			result += getTrailHeads(grid, neighbour, target)
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
