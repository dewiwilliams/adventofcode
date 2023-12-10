package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data [][]int) int64 {
	result := int64(0)

	for _, d := range data {
		result += int64(getPreviousValue(d))
	}

	return result
}
func part1(data [][]int) int64 {
	result := int64(0)

	for _, d := range data {
		result += int64(getNextValue(d))
	}

	return result
}
func getPreviousValue(sequence []int) int {
	if isConstantSequence(sequence) {
		return sequence[0]
	}

	differences := getDifferences(sequence)
	return sequence[0] - getPreviousValue(differences)
}
func getNextValue(sequence []int) int {
	if isConstantSequence(sequence) {
		return sequence[0]
	}

	differences := getDifferences(sequence)
	return sequence[len(sequence)-1] + getNextValue(differences)
}
func isConstantSequence(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] != sequence[i+1] {
			return false
		}
	}

	return true
}
func getDifferences(sequence []int) []int {
	result := []int{}

	for i := 0; i < len(sequence)-1; i++ {
		result = append(result, sequence[i+1]-sequence[i])
	}

	return result
}
func parseData(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := [][]int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		result = append(result, parseIntegers(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseIntegers(line string) []int {
	result := []int{}

	parts := strings.Split(line, " ")
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		value, _ := strconv.ParseInt(p, 10, 32)
		result = append(result, int(value))
	}

	return result
}
