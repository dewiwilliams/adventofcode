package main

import (
	"adventofcode/util"
	"adventofcode/util/grid"
	"fmt"
)

func main() {
	//grid := getData("sample_input.txt")
	grid := getData("input.txt")

	//fmt.Printf("Got grid: %v\n", grid)

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}

func part1(grid grid.Grid) int {
	result := 0

	limit := grid.Width * grid.Height
	for i := range limit {
		if grid.Grid[i] != 1 {
			continue
		}

		neighbours := grid.GetDiagonalNeighbours(i)
		if countCellsWithState(grid, neighbours, 1) < 4 {
			result++
		}
	}

	return result
}
func part2(grid grid.Grid) int {
	result := 0

	for {
		count := removeRolls(grid)
		if count == 0 {
			return result
		}

		result += count
	}
}
func removeRolls(grid grid.Grid) int {
	result := 0

	limit := grid.Width * grid.Height
	for i := range limit {
		if grid.Grid[i] != 1 {
			continue
		}

		neighbours := grid.GetDiagonalNeighbours(i)
		if countCellsWithState(grid, neighbours, 1) >= 4 {
			continue
		}

		grid.Grid[i] = 0
		result++
	}

	return result
}
func countCellsWithState(grid grid.Grid, cells []int, state int) int {
	result := 0

	for _, c := range cells {
		if grid.Grid[c] == state {
			result++
		}
	}

	return result
}

func getData(filename string) grid.Grid {
	lines := util.GetFileLines(filename)

	mapping := make(map[rune]int)
	mapping['.'] = 0
	mapping['@'] = 1

	return grid.NewFromData(lines, mapping)
}
