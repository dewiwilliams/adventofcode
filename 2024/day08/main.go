package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	data, width := getData("input.txt")
	height := len(data) / width
	symbolMap := buildMap(data)

	fmt.Printf("Part 1: %d\n", part1(symbolMap, width, height))
	fmt.Printf("Part 2: %d\n", part2(symbolMap, width, height))
}
func part1(symbolMap map[rune][]int, width, height int) (result int) {
	nodeLocations := make(map[int]bool)

	for _, locations := range symbolMap {
		antiNodes := part1Symbol(locations, width, height)
		for _, v := range antiNodes {
			nodeLocations[v] = true
		}
	}

	return len(nodeLocations)
}
func part1Symbol(locations []int, width, height int) []int {
	result := []int{}

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			antiNodes := getPart1AntinodeLocationCount(locations[i], locations[j], width, height)
			result = append(result, antiNodes...)
		}
	}

	return result
}
func getPart1AntinodeLocationCount(loc1, loc2, width, height int) []int {
	result := []int{}

	x1, y1 := loc1%width, loc1/width
	x2, y2 := loc2%width, loc2/width
	xDiff, yDiff := x2-x1, y2-y1

	antinodeX1, antinodeY1 := x1-xDiff, y1-yDiff
	if inGrid(antinodeX1, antinodeY1, width, height) {
		result = append(result, antinodeX1+antinodeY1*width)
	}

	antinodeX2, antinodeY2 := x2+xDiff, y2+yDiff
	if inGrid(antinodeX2, antinodeY2, width, height) {
		result = append(result, antinodeX2+antinodeY2*width)
	}

	return result
}
func part2(symbolMap map[rune][]int, width, height int) (result int) {
	nodeLocations := make(map[int]bool)

	for _, locations := range symbolMap {
		antiNodes := part2Symbol(locations, width, height)
		for _, v := range antiNodes {
			nodeLocations[v] = true
		}
	}

	return len(nodeLocations)
}
func part2Symbol(locations []int, width, height int) []int {
	result := []int{}

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			antiNodes := getPart2AntinodeLocationCount(locations[i], locations[j], width, height)
			result = append(result, antiNodes...)
		}
	}

	return result
}
func getPart2AntinodeLocationCount(loc1, loc2, width, height int) []int {
	result := []int{}

	x1, y1 := loc1%width, loc1/width
	x2, y2 := loc2%width, loc2/width
	xDiff, yDiff := x2-x1, y2-y1

	for i := 0; ; i++ {
		antinodeX1, antinodeY1 := x1-i*xDiff, y1-i*yDiff
		if inGrid(antinodeX1, antinodeY1, width, height) {
			result = append(result, antinodeX1+antinodeY1*width)
		} else {
			break
		}
	}

	for i := 0; ; i++ {
		antinodeX2, antinodeY2 := x2+i*xDiff, y2+i*yDiff
		if inGrid(antinodeX2, antinodeY2, width, height) {
			result = append(result, antinodeX2+antinodeY2*width)
		} else {
			break
		}
	}

	return result
}
func inGrid(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}
func buildMap(grid []rune) map[rune][]int {
	result := make(map[rune][]int)

	for k, v := range grid {
		if v == '.' {
			continue
		}

		result[v] = append(result[v], k)
	}

	return result
}
func getData(filename string) ([]rune, int) {
	lines := util.GetFileLines(filename)

	width := len(lines[0])
	result := []rune{}

	for _, line := range lines {
		for _, r := range line {
			result = append(result, r)
		}
	}

	return result, width
}
