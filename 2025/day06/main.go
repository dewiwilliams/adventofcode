package main

import (
	"adventofcode/util"
	"adventofcode/util/grid"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sum struct {
	values    []int64
	operation string
}

func main() {
	//filename := "sample_input.txt"
	filename := "input.txt"

	fmt.Printf("Part 1: %d\n", part1(filename))
	fmt.Printf("Part 2: %d\n", part2(filename))
}

func part1(filename string) int64 {
	sums := getPart1Data(filename)

	var result int64

	for _, s := range sums {
		result += part1Sum(s)
	}

	return result
}
func part1Sum(s sum) int64 {
	if s.operation == "+" {
		var result int64

		for _, v := range s.values {
			result += v
		}

		return result
	}

	var result int64
	result = 1

	for _, v := range s.values {
		result *= v
	}

	return result
}

func part2(filename string) int64 {
	grid, operations := getPart2Data(filename)

	emptyColumns := getEmptyColumns(grid)

	if len(emptyColumns)+1 != len(operations) {
		log.Fatalf("Unexpected empty columns: %d vs %d\n", len(emptyColumns)+1, len(operations))
	}

	var result int64

	previousEmptyColumn := -1
	for i, c := range emptyColumns {
		result += part2Value(grid, previousEmptyColumn+1, c, operations[i])
		previousEmptyColumn = c
	}

	result += part2Value(grid, previousEmptyColumn+1, grid.Width, operations[len(operations)-1])

	return result
}

func getEmptyColumns(g grid.Grid) []int {
	result := []int{}

	for x := range g.Width {
		if isColumnEmpty(g, x) {
			result = append(result, x)
		}
	}

	return result
}
func isColumnEmpty(g grid.Grid, x int) bool {
	for y := range g.Height {
		if g.Grid[x+y*g.Width] != -1 {
			return false
		}
	}

	return true
}
func part2Value(g grid.Grid, startColumn, endColumn int, operation string) int64 {
	if operation == "+" {
		var result int64

		for i := startColumn; i < endColumn; i++ {
			result += int64(getValueinColumn(g, i))
		}

		return result
	}

	var result int64
	result = 1

	for i := startColumn; i < endColumn; i++ {
		result *= int64(getValueinColumn(g, i))
	}

	return result
}
func getValueinColumn(g grid.Grid, c int) int {
	var result int
	power := 0

	for i := range g.Height {
		y := g.Height - i - 1
		r := g.Grid[c+y*g.Width]
		if r == -1 && result != 0 {
			return result
		}
		if r != -1 {
			result += util.IntPow(10, power) * r
			power++
		}
	}

	return result
}

func getPart1Data(filename string) []sum {
	lines := util.GetFileLines(filename)

	result := []sum{}

	parts := strings.Fields(lines[0])
	partValues := parseSlice(parts)
	for _, v := range partValues {
		result = append(result, sum{
			values: []int64{v},
		})
	}

	for i := 1; i < len(lines)-1; i++ {
		parts := strings.Fields(lines[i])
		partValues := parseSlice(parts)
		for k, v := range partValues {
			result[k].values = append(result[k].values, v)
		}
	}

	parts = strings.Fields(lines[len(lines)-1])
	for k, v := range parts {
		result[k].operation = v
	}

	return result
}
func getPart2Data(filename string) (grid.Grid, []string) {
	lines := getFileLines(filename)

	mapping := make(map[rune]int)
	grid.GetNumericMapping(mapping)
	mapping[' '] = -1

	g := grid.NewFromData(lines[:len(lines)-1], mapping)
	operations := strings.Fields(lines[len(lines)-1])

	return g, operations
}
func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseSlice(parts []string) []int64 {
	result := []int64{}

	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, int64(n))
	}

	return result
}
