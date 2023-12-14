package main

import (
	"adventofcode/util"
	"fmt"
)

const emptyCell = 0
const squareCell = 1
const roundCell = 2

func main() {
	data, width := parseData("input.txt")

	//printGrid(data, width)
	fmt.Printf("Part 1: %d\n", part1(data, width))
}
func part1(grid []int, width int) int {
	tiltNorth(grid, width)
	//printGrid(grid, width)

	return getTotalLoad(grid, width)
}
func getTotalLoad(grid []int, width int) int {
	height := len(grid) / width

	result := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if grid[cell] != roundCell {
				continue
			}

			result += height - y
		}
	}

	return result
}
func tiltNorth(grid []int, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if grid[cell] != roundCell {
				continue
			}

			targetCell := getNextEmptyCell(grid, cell, -width, y)
			if cell == targetCell {
				continue
			}

			grid[targetCell] = roundCell
			grid[cell] = emptyCell
		}
	}
}
func getNextEmptyCell(grid []int, start, step, maxLength int) int {
	for i := 1; i <= maxLength; i++ {
		if grid[start+i*step] != emptyCell {
			return start + (i-1)*step
		}
	}
	return start + maxLength*step
}
func printGrid(grid []int, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%d", grid[x+y*width])
		}
		fmt.Println()
	}

	fmt.Println()
}
func parseData(filename string) ([]int, int) {
	fileData := util.GetFileLines(filename)

	result := []int{}
	width := 0

	mapping := getRecordMapping()

	for _, line := range fileData {
		if len(line) == 0 {
			continue
		}
		width = len(line)

		parsedLine := util.MapStringToArray(line, mapping)
		result = append(result, parsedLine...)
	}

	return result, width
}
func getRecordMapping() map[rune]int {
	mapping := map[rune]int{}

	mapping['.'] = emptyCell
	mapping['#'] = squareCell
	mapping['O'] = roundCell

	return mapping
}
