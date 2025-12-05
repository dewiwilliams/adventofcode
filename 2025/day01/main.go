package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

func main() {
	data := getData("input")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part1(data []int) int {
	result := 0
	current := 50

	for _, v := range data {
		current += v
		current += 100
		current %= 100

		if current == 0 {
			result++
		}
	}

	return result
}
func part2(data []int) int {
	result := 0
	current := 50

	for _, v := range data {
		previousValue := current

		current += v

		if current >= 100 {
			result += current / 100
		} else if current <= 0 {
			result += -current / 100

			if previousValue > 0 {
				result++
			}
		}

		if current < 0 {
			current += (-current/100 + 1) * 100
		}

		current %= 100
	}

	return result
}

func getData(filename string) []int {
	lines := util.GetFileLines(filename)
	_ = lines

	result := []int{}

	for _, line := range lines {
		value, _ := strconv.ParseInt(line[1:], 10, 32)
		if line[0] == 'L' {
			result = append(result, int(-value))
		} else {
			result = append(result, int(value))
		}

	}

	return result
}
