package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const upDownPipe = 0
const leftRightPipe = 1
const upRightPipe = 2
const rightDownPipe = 3
const downLeftPipe = 4
const leftUpPipe = 5
const emptyCellType = 6
const startCellType = 7

const left = 0
const right = 1
const up = 2
const down = 3

func main() {
	grid, width, start := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(grid, width, start))
	fmt.Printf("Part 2: %d\n", part2(grid, width, start))
}
func part2(gridSource []int, width int, start []int) int {
	grid := make([]int, len(gridSource))
	copy(grid, gridSource)
	startCell := start[0] + start[1]*width

	if grid[startCell] != startCellType {
		panic("Starting cell is not correct")
	}

	grid[startCell] = start[2]
	distanceMap := buildDistanceMap(grid, width, startCell)
	zeroNonLoopCells(grid, distanceMap, startCell)

	return countInternalCells(grid, width)
}
func countInternalCells(grid []int, width int) int {
	result := 0

	for i := range grid {
		if grid[i] != emptyCellType {
			continue
		}
		if getLoopCrossings(grid, width, i)%2 == 1 {
			result++
		}
	}

	return result
}
func zeroNonLoopCells(grid, distanceMap []int, start int) {
	for i := range grid {
		if distanceMap[i] == 0 && i != start {
			grid[i] = emptyCellType
		}
	}
}
func getLoopCrossings(grid []int, width, cell int) int {
	if grid[cell] != emptyCellType {
		panic("Expected empty cell type")
	}

	result := 0

	startX := cell % width
	startY := cell / width

	lineEntryCell := emptyCellType

	for x := startX; x >= 0; x-- {
		currentCell := x + startY*width
		if grid[currentCell] == emptyCellType {
			continue
		} else if grid[currentCell] == upDownPipe {
			result++
		} else if grid[currentCell] == leftUpPipe || grid[currentCell] == downLeftPipe {
			if lineEntryCell != emptyCellType {
				panic("Expected line entry cell type to be empty")
			}
			lineEntryCell = grid[currentCell]
		} else if grid[currentCell] == leftRightPipe {
			if lineEntryCell == emptyCellType {
				panic("Expected line entry cell type to be non-empty")
			}
		} else if grid[currentCell] == rightDownPipe {
			if lineEntryCell == emptyCellType {
				panic("Expected line entry cell type to be non-empty")
			}
			if lineEntryCell == leftUpPipe {
				result++
			}
			lineEntryCell = emptyCellType
		} else if grid[currentCell] == upRightPipe {
			if lineEntryCell == emptyCellType {
				panic("Expected line entry cell type to be non-empty")
			}
			if lineEntryCell == downLeftPipe {
				result++
			}
			lineEntryCell = emptyCellType
		} else {
			panic("Unexpected cell type")
		}
	}

	return result
}
func part1(gridSource []int, width int, start []int) int {
	grid := make([]int, len(gridSource))
	copy(grid, gridSource)

	startCell := start[0] + start[1]*width

	if grid[startCell] != startCellType {
		panic("Starting cell is not correct")
	}

	grid[startCell] = start[2]
	distanceMap := buildDistanceMap(grid, width, startCell)

	return max(distanceMap)
}
func max(grid []int) int {
	result := 0

	for _, v := range grid {
		if v > result {
			result = v
		}
	}

	return result
}
func buildDistanceMap(grid []int, width, start int) []int {
	result := make([]int, len(grid))

	exits := getExitsForType(grid[start])

	followPipe(grid, result, width, start, exits[0])
	followPipe(grid, result, width, start, exits[1])

	return result
}
func followPipe(grid, distanceMap []int, width, start, direction int) {
	previousCell := start
	nextDirection := direction

	for i := 1; ; i++ {
		nextCell := getCellInDirection(previousCell, width, nextDirection)
		if distanceMap[nextCell] != 0 && distanceMap[nextCell] < i {
			return
		} else if nextCell == start {
			return
		}

		distanceMap[nextCell] = i
		exits := getExitsForType(grid[nextCell])
		if getCellInDirection(nextCell, width, exits[0]) == previousCell {
			nextDirection = exits[1]
		} else {
			nextDirection = exits[0]
		}
		previousCell = nextCell
	}
}
func getOppositeDirection(d int) int {
	if d == left {
		return right
	} else if d == right {
		return left
	} else if d == up {
		return down
	} else if d == down {
		return up
	}

	panic("Unknown direction")
}
func getCellInDirection(cell, width, direction int) int {
	if direction == up {
		return cell - width
	} else if direction == down {
		return cell + width
	} else if direction == left {
		return cell - 1
	} else if direction == right {
		return cell + 1
	}

	panic("Unknown direction")
}
func getExitsForType(t int) []int {
	if t == upDownPipe {
		return []int{up, down}
	} else if t == leftRightPipe {
		return []int{left, right}
	} else if t == upRightPipe {
		return []int{up, right}
	} else if t == rightDownPipe {
		return []int{right, down}
	} else if t == downLeftPipe {
		return []int{down, left}
	} else if t == leftUpPipe {
		return []int{left, up}
	}

	fmt.Printf("Unrecognised pipe type: %d\n", t)
	panic("Unrecognised pipe type")
}
func printGrid(grid []int, width int) {
	height := len(grid) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%d", grid[x+y*width])
		}
		fmt.Println("")
	}
}
func parseData(filename string) ([]int, int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	start := parseIntegers(line)

	result := []int{}
	width := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		width = len(line)

		for _, r := range line {
			result = append(result, parseRune(r))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, width, start
}
func parseRune(r rune) int {
	if r == '|' {
		return upDownPipe
	} else if r == '-' {
		return leftRightPipe
	} else if r == 'F' {
		return rightDownPipe
	} else if r == '7' {
		return downLeftPipe
	} else if r == 'J' {
		return leftUpPipe
	} else if r == 'L' {
		return upRightPipe
	} else if r == '.' {
		return emptyCellType
	} else if r == 'S' {
		return startCellType
	}

	panic("Unrecognised value!")
}
func parseIntegers(line string) []int {
	result := []int{}

	parts := strings.Split(line, " ")
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		value, _ := strconv.ParseInt(p, 10, 32)
		result = append(result, int(value))
	}

	return result
}
