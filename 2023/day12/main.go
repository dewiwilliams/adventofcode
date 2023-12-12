package main

import (
	"fmt"
	"strings"
)

const emptyCell = 0
const unknownCell = 1
const filledCell = 2

func main() {
	records, clues := parseData("sample_input.txt")
	fmt.Printf("Got data: %v, %v\n", records, clues)

	fmt.Printf("Part 1: %d\n", part1(records, clues))
	//fmt.Printf("Part 2: %d\n", part2(records, clues))
}
func part1(records, clues [][]int) int {
	result := 0

	for i := range records {
		result += getCombinations(records[i], clues[i])
	}

	return result
}
func part2(records, clues [][]int) int {
	result := 0

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

		result += getCombinations(extendedRecords, extendedClues)

		fmt.Printf("Processed line: %d\n", i)
	}

	return result
}
func getCombinations(records, clues []int) int {
	workspace := make([]int, len(records))
	fillWorkspace(workspace, 0, len(records), emptyCell)

	return getCombinationsImpl(records, clues, 0, 0, workspace)
}
func getCombinationsImpl(records, clues []int, recordOffset, clueOffset int, workspace []int) int {
	if clueOffset == len(clues) {
		if isPatternValid(records, workspace) {
			return 1
		} else {
			return 0
		}
	}
	if recordOffset == len(records) {
		return 0
	}

	result := 0

	for i := recordOffset; i < len(records); i++ {
		if !doesSpringFit(records, i, clues[clueOffset]) {
			continue
		}

		fillWorkspace(workspace, i, clues[clueOffset], filledCell)

		result += getCombinationsImpl(records, clues, i+clues[clueOffset]+1, clueOffset+1, workspace)

		fillWorkspace(workspace, i, clues[clueOffset], emptyCell)
	}

	return result
}
func isPatternValid(records, pattern []int) bool {
	for i := 0; i < len(records); i++ {
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

/*
	func getCombinations(records, clues []int, recordOffset, clueOffset int) int {
		if recordOffset >= len(records) {
			if clueOffset == len(clues) {
				return 1
			}
			return 0
		}
		if clueOffset == len(clues) {
			return 1
		}

		result := 0

		for i := recordOffset; i < len(records); i++ {
			if !doesSpringFit(records, i, clues[clueOffset]) {
				continue
			}

			springLength := clues[clueOffset]
			if i+springLength == len(records) {
				if clueOffset == len(clues)-1 {
					return result + 1
				}
				return result
			}
			if records[i+springLength] == filledCell {
				continue
			}

			result += getCombinations(records, clues, i+springLength+1, clueOffset+1)
		}

		return result
	}
*/
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
	fileData := getFileLines(filename)

	records := [][]int{}
	clues := [][]int{}

	mapping := getRecordMapping()

	for _, line := range fileData {
		parts := strings.Split(line, " ")

		records = append(records, mapStringToArray(parts[0], mapping))
		clues = append(clues, parseIntegerArray(parts[1], ","))
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
