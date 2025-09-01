package main

import (
	"adventofcode/util"
	"fmt"
)

type valueDirection int

const (
	increasing valueDirection = iota
	decreasing
	neutral
)

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data [][]int) int {
	result := 0

	for _, line := range data {
		if part2IsSafe(line) {
			result++
		} else {
		}
	}

	return result
}
func part2IsSafe(data []int) bool {
	if part1IsSafe(data) {
		return true
	}

	for i := range data {
		t := []int{}
		t = append(t, data[:i]...)
		t = append(t, data[i+1:]...)

		if part1IsSafe(t) {
			return true
		}
	}

	return false
}
func part1(data [][]int) int {
	result := 0

	for _, line := range data {
		if part1IsSafe(line) {
			result++
		}
	}

	return result
}
func part1IsSafe(data []int) bool {
	if getValueDirection(data) == neutral {
		return false
	}
	if !isDifferenceSafe(data) {
		return false
	}

	return true
}
func isDifferenceSafe(data []int) bool {
	for i := 1; i < len(data); i++ {
		diff := util.Abs(data[i-1] - data[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
func getValueDirection(data []int) valueDirection {
	direction := getSingleValueDirection(data[0], data[1])
	if direction == neutral {
		return direction
	}

	for i := 2; i < len(data); i++ {
		if direction != getSingleValueDirection(data[i-1], data[i]) {
			return neutral
		}
	}

	return direction
}
func getSingleValueDirection(v1, v2 int) valueDirection {
	if v1 < v2 {
		return increasing
	}
	if v1 > v2 {
		return decreasing
	}
	return neutral
}
func getData(filename string) (result [][]int) {
	lines := util.GetFileLines(filename)

	for _, line := range lines {
		result = append(result, util.ParseIntegerArray(line, " "))
	}

	return
}
