package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func main() {
	data := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data []string) int {
	lensStates := [256][]lens{}

	applyInstructions(&lensStates, data)

	return calculateTotalFocusingPower(&lensStates)
}
func calculateTotalFocusingPower(lensStates *[256][]lens) int {
	result := 0

	for i, state := range lensStates {
		result += (i + 1) * calculateFocusingPower(state)
	}

	return result
}
func calculateFocusingPower(states []lens) int {
	result := 0

	for i, state := range states {
		result += (i + 1) * state.focalLength
	}

	return result
}
func applyInstructions(lensStates *[256][]lens, instructions []string) {
	for _, item := range instructions {
		applyInstruction(lensStates, item)
	}
}
func applyInstruction(lensStates *[256][]lens, instruction string) {
	if strings.Contains(instruction, "=") {
		parts := strings.Split(instruction, "=")

		box := calculateHash(parts[0])
		focalLength, _ := strconv.ParseInt(parts[1], 10, 32)

		existingIndex := findLensIndex(lensStates[box], parts[0])
		if existingIndex == -1 {
			lensStates[box] = append(lensStates[box], lens{
				label:       parts[0],
				focalLength: int(focalLength),
			})
		} else {
			lensStates[box][existingIndex].focalLength = int(focalLength)
		}
	} else {
		parts := strings.Split(instruction, "-")

		box := calculateHash(parts[0])
		lensStates[box] = removeLens(lensStates[box], parts[0])
	}
}
func findLensIndex(lenses []lens, label string) int {
	for i, l := range lenses {
		if l.label == label {
			return i
		}
	}

	return -1
}
func removeLens(lenses []lens, label string) []lens {
	index := findLensIndex(lenses, label)
	if index == -1 {
		return lenses
	}

	return append(lenses[:index], lenses[index+1:]...)
}
func part1(data []string) int {
	result := 0

	for _, item := range data {
		result += calculateHash(item)
	}

	return result
}
func calculateHash(input string) int {
	result := 0

	for _, r := range input {
		result += int(r)
		result *= 17
		result %= 256
	}

	return result
}
func parseData(filename string) []string {
	fileData := util.GetFileLines(filename)

	result := []string{}

	for _, line := range fileData {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ",")
		result = append(result, parts...)
	}

	return result
}
