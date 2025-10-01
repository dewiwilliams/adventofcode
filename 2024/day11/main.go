package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strconv"
)

func main() {
	data := getData("input.txt")
	fmt.Printf("Got data: %v\n", data)

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part1(data []int) int {
	result := 0
	cache := make(map[int]int)

	for _, v := range data {
		result += handleStone(v, 25, cache)
	}

	return result
}
func part2(data []int) int {
	result := 0
	cache := make(map[int]int)

	for _, v := range data {
		result += handleStone(v, 75, cache)
	}

	return result
}
func handleStoneWithMemoisation(stone, remainingIterations int, cache map[int]int) int {
	makeKey := func(stone, remainingIterations int) int {
		return (stone << 8) | remainingIterations
	}

	if stone < 10 {
		key := makeKey(stone, remainingIterations)
		if v, ok := cache[key]; ok {
			return v
		}
	}

	result := handleStone(stone, remainingIterations, cache)

	if stone < 10 {
		key := makeKey(stone, remainingIterations)
		cache[key] = result
	}

	return result
}
func handleStone(stone, remainingIterations int, cache map[int]int) int {
	if remainingIterations == 0 {
		return 1
	}
	if stone == 0 {
		return handleStoneWithMemoisation(1, remainingIterations-1, cache)
	}
	stringValue := strconv.FormatInt(int64(stone), 10)
	stringLength := len(stringValue)
	if stringLength%2 == 0 {
		val1, _ := strconv.ParseInt(stringValue[:stringLength/2], 10, 64)
		val2, _ := strconv.ParseInt(stringValue[stringLength/2:], 10, 64)
		util.RuntimeAssert(val1 < math.MaxInt32)
		util.RuntimeAssert(val2 < math.MaxInt32)

		return handleStoneWithMemoisation(int(val1), remainingIterations-1, cache) + handleStone(int(val2), remainingIterations-1, cache)
	}

	return handleStoneWithMemoisation(stone*2024, remainingIterations-1, cache)
}
func getData(filename string) []int {
	lines := util.GetFileLines(filename)
	util.RuntimeAssert(len(lines) == 1)

	return util.ParseIntegerArray(lines[0], " ")
}
