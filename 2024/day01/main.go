package main

import (
	"adventofcode/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := getData("input.txt")

	part1Data := [2][]int{}
	part1Data[0] = append(part1Data[0], data[0]...)
	part1Data[1] = append(part1Data[1], data[1]...)

	fmt.Printf("Part 1: %d\n", part1(part1Data))
	fmt.Printf("Part 1: %d\n", part2(data))
}
func part1(data [2][]int) int {
	sort.Ints(data[0])
	sort.Ints(data[1])

	result := 0
	for i := range data[0] {
		value1 := data[0][i]
		value2 := data[1][i]
		result += util.Abs(value1 - value2)
	}

	return result
}
func part2(data [2][]int) int {
	count := func(data []int, value int) int {
		result := 0
		for _, v := range data {
			if v == value {
				result++
			}
		}
		return result
	}

	result := 0
	for _, v := range data[0] {
		result += v * count(data[1], v)
	}

	return result
}
func getData(filename string) (result [2][]int) {
	lines := util.GetFileLines(filename)

	for _, line := range lines {
		values := parseIntegers(line)
		result[0] = append(result[0], values[0])
		result[1] = append(result[1], values[1])
	}

	return
}
func parseIntegers(line string) []int {
	result := []int{}

	parts := strings.Split(line, " ")
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		value, _ := strconv.ParseInt(p, 10, 64)
		result = append(result, int(value))
	}

	return result
}
