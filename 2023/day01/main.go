package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data := getData()

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}
func part2(data []string) int {
	result := 0

	for _, v := range data {
		result += extractPart2LineNumber(v)
	}

	return result
}
func extractPart2LineNumber(line string) int {
	firstDigit := ""
	secondDigit := ""

	for i := range line {
		value, _ := extractPart2Number(line, i)
		if value == "" {
			continue
		}

		if firstDigit == "" {
			firstDigit = string(value)
			secondDigit = firstDigit
			continue
		}

		secondDigit = string(value)

		//i += length - 1
	}

	v, _ := strconv.ParseInt(firstDigit+secondDigit, 10, 32)
	return int(v)
}
func extractPart2Number(line string, index int) (string, int) {
	if unicode.IsDigit(rune(line[index])) {
		return string(rune(line[index])), 1
	}

	numbers := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for v, n := range numbers {
		if index+len(n) > len(line) {
			continue
		}
		if n == line[index:index+len(n)] {
			return strconv.Itoa(v), len(n)
		}

	}

	return "", 0
}
func part1(data []string) int {
	result := 0

	for _, v := range data {
		result += extractPart1Number(v)
	}

	return result
}
func extractPart1Number(line string) int {
	firstDigit := ""
	secondDigit := ""

	for _, r := range line {
		if !unicode.IsDigit(r) {
			continue

		}

		if firstDigit == "" {
			firstDigit = string(r)
			secondDigit = firstDigit
			continue
		}

		secondDigit = string(r)
	}

	v, _ := strconv.ParseInt(firstDigit+secondDigit, 10, 32)
	return int(v)
}
func getData() []string {
	result := []string{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			return result
		}

		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
