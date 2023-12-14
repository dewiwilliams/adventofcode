package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

const emptyCell = 0
const squareCell = 1
const roundCell = 2

const north = 0
const east = 1
const south = 2
const west = 3

func main() {
	data, width := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data, width))
	fmt.Printf("Part 2: %d\n", part2(data, width))
}
func part2(grid []int, width int) int {
	start, length := findCycle(grid, width)

	cycleCount := 1000000000
	cycleCount -= start
	cycleCount %= length

	cycleStartGrid := getGridAfterCycles(grid, width, start)
	targetGrid := getGridAfterCycles(cycleStartGrid, width, cycleCount)

	return getTotalLoad(targetGrid, width)
}
func getGridAfterCycles(grid []int, width, cycles int) []int {
	result := make([]int, len(grid))
	copy(result, grid)

	for i := 0; i < cycles; i++ {
		tiltNorth(result, width)
		tiltWest(result, width)
		tiltSouth(result, width)
		tiltEast(result, width)
	}

	return result
}
func findCycle(gridSource []int, width int) (int, int) {
	grid := make([]int, len(gridSource))
	copy(grid, gridSource)

	fingerprints := []string{}

	for i := 0; ; i++ {
		tiltNorth(grid, width)
		tiltWest(grid, width)
		tiltSouth(grid, width)
		tiltEast(grid, width)

		fingerprints = append(fingerprints, getFingerprint(grid))

		start, length := analyseFingerprints(fingerprints)
		if start != -1 && length != -1 {
			return start + 1, length
		}
	}
}
func analyseFingerprints(fingerprints []string) (int, int) {
	if len(fingerprints) < 2 {
		return -1, -1
	}

	for i := 1; i < len(fingerprints); i++ {
		for j := 0; j < i; j++ {
			if fingerprints[i] == fingerprints[j] {
				return j, i - j
			}
		}
	}

	return -1, -1
}
func getFingerprint(grid []int) string {
	parts := []string{}

	for i, v := range grid {
		if v == roundCell {
			parts = append(parts, strconv.Itoa(i))
		}
	}

	return strings.Join(parts, "_")
}
func areGridsEqual(grid1, grid2 []int) bool {
	for i := range grid1 {
		if grid1[i] != grid2[i] {
			return false
		}
	}

	return true
}
func part1(grid []int, width int) int {
	tiltNorth(grid, width)

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
func tiltEast(grid []int, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := width - 1; x >= 0; x-- {
			cell := x + y*width
			if grid[cell] != roundCell {
				continue
			}

			targetCell := getNextEmptyCell(grid, cell, 1, width-x-1)
			if cell == targetCell {
				continue
			}

			grid[targetCell] = roundCell
			grid[cell] = emptyCell
		}
	}
}
func tiltWest(grid []int, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if grid[cell] != roundCell {
				continue
			}

			targetCell := getNextEmptyCell(grid, cell, -1, x)
			if cell == targetCell {
				continue
			}

			grid[targetCell] = roundCell
			grid[cell] = emptyCell
		}
	}
}
func tiltSouth(grid []int, width int) {
	height := len(grid) / width

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if grid[cell] != roundCell {
				continue
			}

			targetCell := getNextEmptyCell(grid, cell, width, height-y-1)
			if cell == targetCell {
				continue
			}

			grid[targetCell] = roundCell
			grid[cell] = emptyCell
		}
	}
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
