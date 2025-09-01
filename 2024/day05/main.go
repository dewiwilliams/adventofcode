package main

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type data struct {
	ordering        map[int][]int
	reverseOrdering map[int][]int

	pageUpdates [][]int
}

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data data) int {
	result := 0

	for _, update := range data.pageUpdates {
		if isOrderingCorrect(data.ordering, update) {
			continue
		}

		correctOrdering := correctOrdering(data.ordering, update)
		middleIndex := (len(correctOrdering) - 1) / 2
		result += correctOrdering[middleIndex]
	}

	return result
}
func correctOrdering(ordering map[int][]int, printing []int) []int {
	duplicate := make([]int, len(printing))
	copy(duplicate, printing)

	slices.SortFunc(duplicate, func(a, b int) int {
		if ordering[a] != nil && slices.Contains(ordering[a], b) {
			return 1
		}
		if ordering[b] != nil && slices.Contains(ordering[b], a) {
			return -1
		}

		return 0
	})

	return duplicate
}
func part1(data data) int {
	result := 0

	for _, update := range data.pageUpdates {
		if !isOrderingCorrect(data.ordering, update) {
			continue
		}

		middleIndex := (len(update) - 1) / 2
		result += update[middleIndex]
	}

	return result
}
func isOrderingCorrect(ordering map[int][]int, printing []int) bool {
	for k, v := range printing {
		mustAppearBefore := ordering[v]
		if mustAppearBefore == nil {
			continue
		}

		for _, v2 := range mustAppearBefore {
			if slices.Contains(printing[:k], v2) {
				return false
			}
		}
	}

	return true
}
func getData(filename string) data {
	lines := util.GetFileLines(filename)

	result := data{
		ordering:        make(map[int][]int),
		reverseOrdering: make(map[int][]int),
	}
	emptyLine := -1

	for i := range len(lines) {
		if len(lines[i]) == 0 {
			emptyLine = i
			break
		}

		parts := strings.Split(lines[i], "|")
		val1, _ := strconv.ParseInt(parts[0], 10, 32)
		val1Int := int(val1)
		val2, _ := strconv.ParseInt(parts[1], 10, 32)
		val2Int := int(val2)
		result.ordering[val1Int] = append(result.ordering[val1Int], int(val2))
		result.reverseOrdering[val2Int] = append(result.ordering[val2Int], int(val1))
	}

	for i := emptyLine + 1; i < len(lines); i++ {
		result.pageUpdates = append(result.pageUpdates, util.ParseIntegerArray(lines[i], ","))
	}

	return result
}
