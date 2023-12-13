package main

import (
	"adventofcode/util"
	"fmt"
)

const ashCell = 0
const rocksCell = 1

type grid struct {
	data   []int
	width  int
	height int
}

func main() {
	grids := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(grids))
}
func part1(grids []grid) int {
	result := 0

	for _, g := range grids {
		result += getPart1GridResult(g)
	}

	return result
}
func getPart1GridResult(grid grid) int {
	for i := 0; i < grid.width; i++ {
		if isVerticalMirrorLine(grid, i) {
			return i + 1
		}
	}

	for i := 0; i < grid.height; i++ {
		if isHorizontalMirrorLine(grid, i) {
			return (i + 1) * 100
		}
	}

	panic("No line of symmetry found")
}
func isVerticalMirrorLine(grid grid, x int) bool {
	for i := 0; ; i++ {
		start1 := x - i
		start2 := x + i + 1

		if start1 < 0 || start2 >= grid.width {
			return i != 0
		}

		if !doLinesMatch(grid, start1, start2, grid.width, grid.height) {
			return false
		}
	}
}
func isHorizontalMirrorLine(grid grid, y int) bool {
	for i := 0; ; i++ {
		start1 := (y - i) * grid.width
		start2 := (y + i + 1) * grid.width

		if y-i < 0 || y+i+1 >= grid.height {
			return i != 0
		}

		if !doLinesMatch(grid, start1, start2, 1, grid.width) {
			return false
		}
	}
}
func doLinesMatch(grid grid, start1, start2, step, length int) bool {
	for i := 0; i < length; i++ {
		cell1 := start1 + i*step
		cell2 := start2 + i*step

		if grid.data[cell1] != grid.data[cell2] {
			return false
		}
	}
	return true
}
func printGrid(grid grid) {
	height := len(grid.data) / grid.width

	for y := 0; y < height; y++ {
		for x := 0; x < grid.width; x++ {
			fmt.Printf("%d", grid.data[x+y*grid.width])
		}
		fmt.Println()
	}
}
func parseData(filename string) []grid {
	fileData := util.GetFileLines(filename)

	result := []grid{}

	currentGrid := []int{}
	currentWidth := 0

	mapping := getRecordMapping()

	for _, line := range fileData {
		if len(line) == 0 {

			if len(currentGrid) != 0 {
				result = append(result, grid{
					data:   currentGrid,
					width:  currentWidth,
					height: len(currentGrid) / currentWidth,
				})

				currentGrid = []int{}
				currentWidth = 0
			}

			continue
		}

		parsedLine := util.MapStringToArray(line, mapping)
		currentGrid = append(currentGrid, parsedLine...)
		currentWidth = len(line)
	}

	if len(currentGrid) != 0 {
		result = append(result, grid{
			data:   currentGrid,
			width:  currentWidth,
			height: len(currentGrid) / currentWidth,
		})
	}

	return result
}
func getRecordMapping() map[rune]int {
	mapping := map[rune]int{}

	mapping['.'] = ashCell
	mapping['#'] = rocksCell

	return mapping
}
