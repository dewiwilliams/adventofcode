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
	fmt.Printf("Part 2: %d\n", part2(grids))
}
func part2(grids []grid) int {
	result := 0

	for _, g := range grids {
		result += getGridResult(g, 1)
	}

	return result
}
func part1(grids []grid) int {
	result := 0

	for _, g := range grids {
		result += getGridResult(g, 0)
	}

	return result
}
func getGridResult(grid grid, requiredDifference int) int {
	for i := 0; i < grid.width; i++ {
		if getVerticalMirrorLineDifference(grid, i) == requiredDifference {
			return i + 1
		}
	}

	for i := 0; i < grid.height; i++ {
		if getHorizontalMirrorLineDifference(grid, i) == requiredDifference {
			return (i + 1) * 100
		}
	}

	panic("No line of symmetry found")
}

func getVerticalMirrorLineDifference(grid grid, x int) int {
	result := 0

	for i := 0; ; i++ {
		start1 := x - i
		start2 := x + i + 1

		if start1 < 0 || start2 >= grid.width {
			if i == 0 {
				return 1000
			}

			return result
		}

		result += getLineDifference(grid, start1, start2, grid.width, grid.height)
	}
}

func getHorizontalMirrorLineDifference(grid grid, y int) int {
	result := 0

	for i := 0; ; i++ {
		start1 := (y - i) * grid.width
		start2 := (y + i + 1) * grid.width

		if y-i < 0 || y+i+1 >= grid.height {
			if i == 0 {
				return 1000
			}

			return result
		}

		result += getLineDifference(grid, start1, start2, 1, grid.width)
	}
}
func getLineDifference(grid grid, start1, start2, step, length int) int {
	result := 0

	for i := 0; i < length; i++ {
		cell1 := start1 + i*step
		cell2 := start2 + i*step

		if grid.data[cell1] != grid.data[cell2] {
			result++
		}
	}
	return result
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
