package main

import (
	"adventofcode/util"
	"adventofcode/util/grid"
	"fmt"
)

const cellTypeEmpty = 0
const cellTypeSplitter = 1
const cellTypeStart = 2
const cellTypeBeam = 3

func main() {
	//filename := "sample_input.txt"
	filename := "input.txt"

	fmt.Printf("Part 1: %d\n", part1(filename))
	fmt.Printf("Part 2: %d\n", part2(filename))
}
func part1(filename string) int {
	grid := getData(filename)

	startCells := grid.GetCellsWithValue(2)
	util.RuntimeAssert(len(startCells) == 1)
	start := startCells[0]

	result := 0

	grid.Grid[start+grid.Width] = cellTypeBeam

	for y := 0; y < grid.Height-1; y++ {
		result += splitRowBeam(grid, y)
	}

	return result
}
func splitRowBeam(g grid.Grid, row int) int {
	cells := g.GetCellsWithValueInRow(row, cellTypeBeam)

	result := 0

	for _, cell := range cells {
		nextCell := cell + g.Width

		if g.Grid[nextCell] == 1 {
			g.Grid[nextCell-1] = cellTypeBeam
			g.Grid[nextCell+1] = cellTypeBeam

			result++
		} else {
			g.Grid[nextCell] = cellTypeBeam
		}
	}

	return result
}

func part2(filename string) int {
	grid := getData(filename)

	lookupMap := make([]int, grid.Width*grid.Height)

	bottomRowCells := grid.GetCellsInRow(grid.Height - 1)
	for _, cell := range bottomRowCells {
		lookupMap[cell] = 1
	}

	for y := grid.Height - 2; y >= 0; y-- {
		buildLookupForRow(grid, y, lookupMap)
	}

	startCells := grid.GetCellsWithValue(2)
	util.RuntimeAssert(len(startCells) == 1)
	start := startCells[0]

	return lookupMap[start]
}
func buildLookupForRow(grid grid.Grid, row int, lookupMap []int) {
	for x := range grid.Width {
		cell := x + row*grid.Width
		cellType := grid.Grid[cell]

		if cellType == cellTypeEmpty || cellType == cellTypeStart {
			lookupMap[cell] = lookupMap[cell+grid.Width]
		} else if cellType == cellTypeSplitter {
			lookupMap[cell] = lookupMap[cell+grid.Width+1] + lookupMap[cell+grid.Width-1]
		} else {
			util.RuntimeAssert(false)
		}
	}
}

func getData(filename string) grid.Grid {
	lines := util.GetFileLines(filename)

	mapping := make(map[rune]int)
	mapping['.'] = cellTypeEmpty
	mapping['^'] = cellTypeSplitter
	mapping['S'] = cellTypeStart
	mapping['|'] = cellTypeBeam

	return grid.NewFromData(lines, mapping)
}
