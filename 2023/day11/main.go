package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	grid, width := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(grid, width))
	fmt.Printf("Part 2: %d\n", part2(grid, width))
}
func part1(grid []bool, width int) int64 {
	return calculateGalaxyDistances(grid, width, 1)
}
func part2(grid []bool, width int) int64 {
	return calculateGalaxyDistances(grid, width, 1000000-1)
}
func calculateGalaxyDistances(grid []bool, width, emptySpaceGap int) int64 {
	emptyRows := getEmptyRows(grid, width)
	emptyColumns := getEmptyColumns(grid, width)

	result := int64(0)
	galaxies := getGalaxies(grid)
	galaxyCount := len(galaxies)

	for i := 0; i < galaxyCount; i++ {
		for j := i + 1; j < galaxyCount; j++ {
			xi := galaxies[i] % width
			yi := galaxies[i] / width
			xj := galaxies[j] % width
			yj := galaxies[j] / width

			thisDistance := getManhattanDistance(xi, yi, xj, yj)
			thisDistance += emptySpaceGap * getValueRangeInArray(xi, xj, emptyColumns)
			thisDistance += emptySpaceGap * getValueRangeInArray(yi, yj, emptyRows)

			result += int64(thisDistance)
		}
	}

	return result
}
func getValueRangeInArray(v1, v2 int, values []int) int {
	result := 0

	r := []int{v1, v2}
	sort.Ints(r)

	for i := r[0]; i < r[1]; i++ {
		if slices.Contains(values, i) {
			result++
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
