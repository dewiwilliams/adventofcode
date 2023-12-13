package main

import (
	"fmt"
	"strings"

	"adventofcode/util"
)

const emptyCell = 0
const unknownCell = 1
const filledCell = 2

func main() {
	records, clues := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(records, clues))
	fmt.Printf("Part 2: %d\n", part2(records, clues))
}
func part1(records, clues [][]int) int {
	result := 0

	for i := range records {
		result += getCombinations(records[i], clues[i])
	}

	return result
}
func part2(records, clues [][]int) int64 {
	result := int64(0)

	for i := range records {
		extendedRecords := []int{}
		extendedClues := []int{}

		for j := 0; j < 5; j++ {
			extendedRecords = append(extendedRecords, records[i]...)
			if j < 4 {
				extendedRecords = append(extendedRecords, unknownCell)
			}
			extendedClues = append(extendedClues, clues[i]...)
		}

		result += int64(getCombinations(extendedRecords, extendedClues))
	}

	return result
}
func getCombinations(records, clues []int) int {
	workspace := make([]int, len(records))
	fillWorkspace(workspace, 0, len(records), emptyCell)

	cache := make(map[string]int)

	return getCombinationsImpl(records, clues, 0, 0, workspace, cache)
}
func getCombinationsImpl(records, clues []int, recordOffset, clueOffset int, workspace []int, cache map[string]int) int {
	result := 0

	for i := recordOffset; i < len(records); i++ {
		if records[i] == emptyCell {
			continue
		}
		if !doesSpringFit(records, i, clues[clueOffset]) {
			continue
		}

		fillWorkspace(workspace, i, clues[clueOffset], filledCell)
		if isPatternValid(records, workspace, util.Min(len(workspace), i+clues[clueOffset])) {
			result += getCachedCombinations(records, clues, i+clues[clueOffset]+1, clueOffset+1, workspace, cache)
		}
		fillWorkspace(workspace, i, clues[clueOffset], emptyCell)
	}

	return result
}
func getCachedCombinations(records, clues []int, recordOffset, clueOffset int, workspace []int, cache map[string]int) int {
	if clueOffset == len(clues) {
		if isPatternValid(records, workspace, len(records)) {
			return 1
		} else {
			return 0
		}
	}
	if recordOffset >= len(records) {
		return 0
	}

	key := fmt.Sprintf("%d_%d", recordOffset, clueOffset)
	if val, ok := cache[key]; ok {
		return val
	}

	val := getCombinationsImpl(records, clues, recordOffset, clueOffset, workspace, cache)
	cache[key] = val
	return val
}
func isPatternValid(records, pattern []int, length int) bool {
	for i := 0; i < length; i++ {
		if pattern[i] == filledCell && records[i] == emptyCell {
			return false
		} else if pattern[i] == emptyCell && records[i] == filledCell {
			return false
		}
	}

	return true
}
func fillWorkspace(target []int, start, length, value int) {
	for i := 0; i < length; i++ {
		target[start+i] = value
	}
}
func doesSpringFit(records []int, recordOffset, clue int) bool {
	if recordOffset+clue > len(records) {
		return false
	}

	for i := recordOffset; i < clue; i++ {
		if records[i] == emptyCell {
			return false
		}
	}

	return true
}
func parseData(filename string) ([][]int, [][]int) {
	fileData := util.GetFileLines(filename)

	records := [][]int{}
	clues := [][]int{}

	mapping := getRecordMapping()

	for _, line := range fileData {
		parts := strings.Split(line, " ")

		records = append(records, util.MapStringToArray(parts[0], mapping))
		clues = append(clues, util.ParseIntegerArray(parts[1], ","))
	}

	return records, clues
}
func getRecordMapping() map[rune]int {
	mapping := map[rune]int{}

	mapping['.'] = emptyCell
	mapping['?'] = unknownCell
	mapping['#'] = filledCell

	return mapping
}
