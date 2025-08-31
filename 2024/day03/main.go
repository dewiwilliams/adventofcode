package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))

}
func part1(data []string) int {
	r, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v\n", err)
	}

	result := 0
	for _, line := range data {
		matches := r.FindAll([]byte(line), -1)
		for _, match := range matches {
			result += mul(string(match))
		}
	}

	return result
}
func part2(data []string) int {
	r, err := regexp.Compile(`mul\(\d+,\d+\)`)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v\n", err)
	}

	allData := ""
	for _, line := range data {
		allData += line
	}

	result := 0

	lineString := []byte(allData)
	matches := r.FindAll(lineString, -1)
	matchesIndexes := r.FindAllIndex(lineString, -1)
	doIndexes := findAllIndex(allData, "do()")
	dontIndexes := findAllIndex(allData, "don't()")

	for i, matchIndex := range matchesIndexes {
		if !isEnabled(matchIndex[0], doIndexes, dontIndexes) {
			continue
		}
		result += mul(string(matches[i]))
	}

	return result
}
func isEnabled(index int, doIndex, dontIndex []int) bool {
	lastDoIndex := getLastIndex(index, doIndex)
	lastDontIndex := getLastIndex(index, dontIndex)

	if lastDoIndex == -1 && lastDontIndex == -1 {
		return true
	}

	return lastDoIndex > lastDontIndex
}
func getLastIndex(index int, list []int) int {
	result := -1

	for _, v := range list {
		if index < v {
			return result
		}
		result = v
	}

	return result
}
func findAllIndex(s, substr string) []int {
	result := []int{}
	lastIndex := 0

	for {
		target := s[lastIndex:]
		index := strings.Index(target, substr)
		if index == -1 {
			return result
		}
		result = append(result, lastIndex+index)
		lastIndex += index + 1
	}
}
func mul(command string) int {
	commaIndex := strings.Index(command, ",")
	first := command[4:commaIndex]
	second := command[commaIndex+1 : len(command)-1]

	firstValue, err := strconv.ParseInt(first, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse: %v\n", first)
	}
	secondValue, err := strconv.ParseInt(second, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse: %v\n", first)
	}

	return int(firstValue * secondValue)
}
func getData(filename string) []string {
	return util.GetFileLines(filename)
}
