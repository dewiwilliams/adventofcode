package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	grid, width := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(grid, width))
	fmt.Printf("Part 2: %d\n", part2(grid, width))
}
func part2(data []rune, width int) int {
	result := 0
	height := len(data) / width
	numberMap := buildNumberMap(data, width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if data[cell] == '*' {
				result += getGearRatio(numberMap, width, x, y)
			}
		}
	}

	return result
}
func getGearRatio(numberMap []int, width, x, y int) int {
	cell := x + y*width

	// I've expanded the grid so that I don't need to think about edge case on the boundary... easy!

	partNumbers := []int{}
	partNumbers = append(partNumbers, getPartNumbersFromSingleCell(numberMap, cell-1)...)
	partNumbers = append(partNumbers, getPartNumbersFromSingleCell(numberMap, cell+1)...)
	partNumbers = append(partNumbers, getPartNumbersFromVerticalBaseCell(numberMap, cell-width)...)
	partNumbers = append(partNumbers, getPartNumbersFromVerticalBaseCell(numberMap, cell+width)...)

	if len(partNumbers) == 2 {
		return partNumbers[0] * partNumbers[1]
	}

	return 0
}
func getPartNumbersFromSingleCell(numberMap []int, cell int) []int {
	if numberMap[cell] == 0 {
		return []int{}
	} else {
		return []int{numberMap[cell]}
	}
}
func getPartNumbersFromVerticalBaseCell(numberMap []int, baseCell int) []int {
	result := []int{}

	if numberMap[baseCell] == 0 {
		if numberMap[baseCell-1] != 0 {
			result = append(result, numberMap[baseCell-1])
		}
		if numberMap[baseCell+1] != 0 {
			result = append(result, numberMap[baseCell+1])
		}
	} else {
		result = append(result, numberMap[baseCell])
	}

	return result
}
func buildNumberMap(data []rune, width int) []int {
	result := make([]int, len(data))
	height := len(data) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if !unicode.IsDigit(data[cell]) {
				continue
			}

			length := getNumberLength(data, width, x, y)
			value := getNumber(data, cell, length)

			for i := 0; i < length; i++ {
				result[cell+i] = value
			}

			x += length
		}
	}

	return result
}
func part1(data []rune, width int) int {
	adjacencyMap := buildAdjacencyMap(data, width, [2]rune{'*', '+'})

	result := 0
	height := len(data) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if !unicode.IsDigit(data[cell]) {
				continue
			}

			length := getNumberLength(data, width, x, y)
			if !isAdjacentToSymbol(adjacencyMap, cell, length) {
				continue
			}

			result += getNumber(data, cell, length)
			x += length
		}
	}

	return result
}
func getNumberLength(data []rune, width, x, y int) int {
	for i := x + 1; i < width; i++ {
		cell := i + y*width

		if !unicode.IsDigit(data[cell]) {
			return i - x
		}
	}

	return width - x
}
func getNumber(data []rune, start, length int) int {
	result := ""

	for i := 0; i < length; i++ {
		result += string(data[start+i])
	}

	r, _ := strconv.ParseInt(result, 10, 32)
	return int(r)
}
func isAdjacentToSymbol(adjacencyMap []bool, start, length int) bool {
	for i := 0; i < length; i++ {
		if adjacencyMap[start+i] {
			return true
		}
	}

	return false
}
func buildAdjacencyMap(data []rune, width int, symbols [2]rune) []bool {
	result := make([]bool, len(data))
	height := len(data) / width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := x + y*width

			if data[cell] == symbols[0] || data[cell] == symbols[1] {
				neighbours := getNeighbours(x, y, width, height)

				for _, n := range neighbours {
					result[n] = true
				}
			}
		}
	}

	return result
}
func getNeighbours(x, y, width, height int) []int {
	result := []int{}

	if x > 0 {
		result = append(result, (x-1)+y*width)
	}
	if x < width-1 {
		result = append(result, (x+1)+y*width)
	}
	if y > 0 {
		result = append(result, x+(y-1)*width)
	}
	if y < height-1 {
		result = append(result, x+(y+1)*width)
	}
	if x > 0 && y > 0 {
		result = append(result, (x-1)+(y-1)*width)
	}
	if x < width-1 && y > 0 {
		result = append(result, (x+1)+(y-1)*width)
	}
	if x > 0 && y < height-1 {
		result = append(result, (x-1)+(y+1)*width)
	}
	if x < width-1 && y < height-1 {
		result = append(result, (x+1)+(y+1)*width)
	}

	return result
}
func getData(filename string) ([]rune, int) {
	result := []rune{}
	width := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		width = len(line)

		if lineNumber == 0 {
			result = append(result, makeEmptyLine(width+2)...)
		}

		result = append(result, '.')
		for _, r := range line {
			if unicode.IsDigit(r) {
				result = append(result, r)
			} else if r == '.' {
				result = append(result, r)
			} else if r == '*' {
				result = append(result, '*')
			} else {
				result = append(result, '+')
			}
		}
		result = append(result, '.')

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result = append(result, makeEmptyLine(width+2)...)

	return result, width + 2
}
func makeEmptyLine(length int) []rune {
	result := make([]rune, length)

	for i := 0; i < length; i++ {
		result[i] = '.'
	}

	return result
}
