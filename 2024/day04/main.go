package main

import (
	"adventofcode/util"
	"fmt"
)

func main() {
	grid, width := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(grid, width))
	fmt.Printf("Part 2: %d\n", part2(grid, width))
}
func part2(data []rune, width int) int {
	height := len(data) / width

	result := 0

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if isCrossMas(data, x, y, width) {
				result++
			}
		}
	}

	return result

}
func isCrossMas(data []rune, x, y, width int) bool {
	base := x + y*width

	if data[base] != 'A' {
		return false
	}

	negative1 := data[base-width-1] == 'M' && data[base+width+1] == 'S'
	negative2 := data[base-width-1] == 'S' && data[base+width+1] == 'M'
	if !negative1 && !negative2 {
		return false
	}

	positive1 := data[base+width-1] == 'M' && data[base-width+1] == 'S'
	positive2 := data[base+width-1] == 'S' && data[base-width+1] == 'M'
	if !positive1 && !positive2 {
		return false
	}

	return true
}
func part1(data []rune, width int) int {
	height := len(data) / width

	result := 0

	for y := range height {
		for x := range width {
			start := x + y*width

			strides := getStrides(x, y, width, height)
			for _, stride := range strides {
				if isXMAS(data, start, stride) {
					result++
				}
			}
		}
	}

	return result
}
func getStrides(x, y, width, height int) []int {
	result := []int{}

	if x >= 3 {
		result = append(result, -1)
	}
	if x < width-3 {
		result = append(result, 1)
	}
	if y >= 3 {
		result = append(result, -width)
	}
	if y < height-3 {
		result = append(result, width)
	}
	if x >= 3 && y >= 3 {
		result = append(result, -width-1)
	}
	if x >= 3 && y < height-3 {
		result = append(result, width-1)
	}
	if x < width-3 && y >= 3 {
		result = append(result, -width+1)
	}
	if x < width-3 && y < height-3 {
		result = append(result, width+1)
	}

	return result
}
func isXMAS(data []rune, start, stride int) bool {
	return data[start+0*stride] == 'X' &&
		data[start+1*stride] == 'M' &&
		data[start+2*stride] == 'A' &&
		data[start+3*stride] == 'S'
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
