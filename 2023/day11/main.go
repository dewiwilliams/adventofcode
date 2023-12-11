package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid, width := parseData("input.txt")
	printGrid(grid, width)

	fmt.Println()

	fmt.Printf("Part 1: %d\n", part1(grid, width))
}
func part1(gridSource []bool, widthSource int) int {
	grid, width := expandGrid(gridSource, widthSource)
	printGrid(grid, width)

	result := 0
	galaxies := getGalaxies(grid)
	galaxyCount := len(galaxies)

	for i := 0; i < galaxyCount; i++ {
		for j := i + 1; j < galaxyCount; j++ {
			xi := galaxies[i] % width
			yi := galaxies[i] / width
			xj := galaxies[j] % width
			yj := galaxies[j] / width

			result += getManhattanDistance(xi, yi, xj, yj)
		}
	}

	return result
}
func getManhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}
func abs(v int) int {
	if v >= 0 {
		return v
	}

	return -v
}
func getGalaxies(grid []bool) []int {
	result := []int{}

	for i := range grid {
		if grid[i] {
			result = append(result, i)
		}
	}

	return result
}
func expandGrid(grid []bool, width int) ([]bool, int) {
	grid2, width2 := doubleEmptyColumns(grid, width)
	return doubleEmptyRows(grid2, width2)
}
func doubleEmptyRows(grid []bool, width int) ([]bool, int) {
	height := len(grid) / width
	emptyRows := getEmptyRows(grid, width)
	emptyIndex := 0
	newRows := [][]bool{}

	for y := 0; y < height; y++ {
		row := getRow(grid, width, y)
		newRows = append(newRows, row)

		if emptyIndex < len(emptyRows) && emptyRows[emptyIndex] == y {
			newRows = append(newRows, row)
			emptyIndex++
		}
	}

	return mergeRows(newRows), width
}
func doubleEmptyColumns(grid []bool, width int) ([]bool, int) {
	emptyColumns := getEmptyColumns(grid, width)
	emptyIndex := 0
	newColumns := [][]bool{}

	for x := 0; x < width; x++ {
		column := getColumn(grid, width, x)
		newColumns = append(newColumns, column)

		if emptyIndex < len(emptyColumns) && emptyColumns[emptyIndex] == x {
			newColumns = append(newColumns, column)
			emptyIndex++
		}
	}

	return mergeColumns(newColumns), width + len(emptyColumns)
}
func mergeColumns(columns [][]bool) []bool {
	width := len(columns)
	height := len(columns[0])
	result := make([]bool, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			result[cell] = columns[x][y]
		}
	}

	return result
}
func getColumn(grid []bool, width, column int) []bool {
	height := len(grid) / width
	result := make([]bool, height)
	for y := 0; y < height; y++ {
		result[y] = grid[column+y*width]
	}
	return result
}
func mergeRows(rows [][]bool) []bool {
	result := rows[0]
	for i := 1; i < len(rows); i++ {
		result = append(result, rows[i]...)
	}
	return result
}
func getRow(grid []bool, width, row int) []bool {
	result := make([]bool, width)
	for x := 0; x < width; x++ {
		result[x] = grid[x+row*width]
	}
	return result
}
func getEmptyColumns(grid []bool, width int) []int {
	result := []int{}
	height := len(grid) / width

	for x := 0; x < width; x++ {
		if areCellsEmpty(grid, x, width, height) {
			result = append(result, x)
		}
	}

	return result
}
func getEmptyRows(grid []bool, width int) []int {
	result := []int{}
	height := len(grid) / width

	for y := 0; y < height; y++ {
		if areCellsEmpty(grid, y*width, 1, width) {
			result = append(result, y)
		}
	}

	return result
}
func areCellsEmpty(grid []bool, start, step, count int) bool {
	for i := 0; i < count; i++ {
		if grid[start+i*step] {
			return false
		}
	}

	return true
}
func printGrid(grid []bool, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width
			if grid[cell] {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}
func parseData(filename string) ([]bool, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []bool{}
	width := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		width = len(line)

		for _, r := range line {
			if r == '#' {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, width
}
