package main

import (
	"adventofcode/util"
	"fmt"
)

type cellType int

const (
	empty cellType = iota
	obstacle
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func main() {
	grid, width := getData("input.txt")
	translatedGrid, startingPoint := translateGrid(grid)

	fmt.Printf("Grid size: %d\n", len(translatedGrid))

	fmt.Printf("Part 1: %d\n", part1(translatedGrid, width, startingPoint))
	fmt.Printf("Part 2: %d\n", part2(translatedGrid, width, startingPoint))
}
func part2(grid []cellType, width int, startingPoint int) int {
	result := 0

	for i := range len(grid) {
		if grid[i] != empty {
			continue
		}

		grid[i] = obstacle
		if loops(grid, width, startingPoint) {
			result++
		}
		grid[i] = empty
	}

	return result
}
func loops(grid []cellType, width int, startingPoint int) bool {
	currentX := startingPoint % width
	currentY := startingPoint / width
	height := len(grid) / width
	currentDirection := up
	statesSeen := make(map[int]bool)

	for {
		currentCell := currentX + currentY*width
		key := int(currentDirection) | (currentCell << 8)
		if statesSeen[key] {
			return true
		}
		statesSeen[key] = true

		nextX, nextY, nextDirection := getNextPoint(currentX, currentY, currentDirection, grid, width, height)
		if nextX == -1 || nextY == -1 {
			return false
		}

		currentX, currentY, currentDirection = nextX, nextY, nextDirection
	}
}
func part1(grid []cellType, width int, startingPoint int) int {
	currentX := startingPoint % width
	currentY := startingPoint / width
	height := len(grid) / width
	currentDirection := up
	coverage := make([]bool, len(grid))

	for {
		currentPoint := currentX + currentY*width
		coverage[currentPoint] = true

		nextX, nextY, nextDirection := getNextPoint(currentX, currentY, currentDirection, grid, width, height)
		if nextX == -1 || nextY == -1 {
			return getCoverage(coverage)
		}

		currentX, currentY, currentDirection = nextX, nextY, nextDirection
	}
}
func getNextPoint(x, y int, direction direction, grid []cellType, width, height int) (int, int, direction) {
	nextX, nextY := x, y

	if direction == up {
		nextY--
	} else if direction == down {
		nextY++
	} else if direction == left {
		nextX--
	} else if direction == right {
		nextX++
	}

	if nextX < 0 || nextY < 0 || nextX >= width || nextY >= height {
		return -1, -1, up
	}

	currentCell := nextX + nextY*width
	if grid[currentCell] == empty {
		return nextX, nextY, direction
	}

	direction++
	direction %= 4

	return x, y, direction
}
func getCoverage(coverage []bool) int {
	result := 0

	for _, v := range coverage {
		if v {
			result++
		}
	}

	return result
}
func translateGrid(grid []rune) ([]cellType, int) {
	result := make([]cellType, len(grid))
	startingPosition := -1

	for k, v := range grid {
		if v == '.' {
			result[k] = empty
		} else if v == '#' {
			result[k] = obstacle
		} else if v == '^' {
			startingPosition = k
		}
	}

	return result, startingPosition
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
