package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

type sum struct {
	total  int64
	values []int64
}
type operator int

const (
	add operator = iota
	multiply
	concatenate
)

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(sums []sum) int {
	result := 0

	for _, s := range sums {
		if isSatisfiablePart2(s) {
			result += int(s.total)
		}
	}

	return result
}
func isSatisfiablePart2(data sum) bool {
	limit := 1
	for range len(data.values) - 1 {
		limit *= 3
	}

	for i := range limit {
		if isSatisfiedPart2(data, i) {
			return true
		}

	}

	return false
}
func isSatisfiedPart2(data sum, operators int) bool {
	result := data.values[0]

	for i := range len(data.values) - 1 {
		operator := operator(operators % 3)
		if operator == add {
			result += data.values[i+1]
		} else if operator == multiply {
			result *= data.values[i+1]
		} else if operator == concatenate {
			stringValue := fmt.Sprintf("%d%d", result, data.values[i+1])
			result, _ = strconv.ParseInt(stringValue, 10, 64)
		}

		operators /= 3
	}

	return result == data.total
}
func part1(sums []sum) int {
	result := 0

	for _, s := range sums {
		if isSatisfiablePart1(s) {
			result += int(s.total)
		}
	}

	return result
}
func isSatisfiablePart1(data sum) bool {
	limit := 1<<len(data.values) - 1

	for i := range limit {
		if isSatisfiedPart1(data, i) {
			return true
		}

	}

	return false
}
func isSatisfiedPart1(data sum, operators int) bool {
	result := data.values[0]

	for i := range len(data.values) - 1 {
		if operator(operators&0x1) == add {
			result += data.values[i+1]
		} else {
			result *= data.values[i+1]
		}

		operators >>= 1
	}

	return result == data.total
}
func getData(filename string) []sum {
	result := []sum{}

	lines := util.GetFileLines(filename)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		total, _ := strconv.ParseInt(parts[0][:len(parts[0])-1], 10, 64)
		lineData := sum{
			total: total,
		}

		for _, v := range parts[1:] {
			value, _ := strconv.ParseInt(v, 10, 64)
			lineData.values = append(lineData.values, value)
		}

		result = append(result, lineData)
	}

	return result
}
