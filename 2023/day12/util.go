package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := []string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func mapStringToArray(source string, mapping map[rune]int) []int {
	result := make([]int, len(source))

	for i, r := range source {
		result[i] = mapping[r]
	}

	return result
}
func parseIntegerArray(source, separator string) []int {
	result := []int{}

	parts := strings.Split(source, separator)
	for _, p := range parts {
		val, _ := strconv.ParseInt(p, 10, 32)
		result = append(result, int(val))
	}

	return result
}
