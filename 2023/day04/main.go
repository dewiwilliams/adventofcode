package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	data := getData("input.txt")

	fmt.Printf("Part1: %v\n", part1(data))
	fmt.Printf("Part2: %v\n", part2(data))
}
func part2(data [][]int) int {
	result := 0
	cardCount := len(data) / 2

	cardCounts := make([]int, cardCount)
	for i := 0; i < cardCount; i++ {
		cardCounts[i] = 1
	}

	for i := 0; i < cardCount; i++ {
		matchingNumbers := findMatchingNumbers(data[i*2], data[i*2+1])
		for j := 1; j <= matchingNumbers; j++ {
			if i+j < cardCount {
				cardCounts[i+j] += cardCounts[i]
			}
		}

		result += cardCounts[i]
	}

	return result
}
func part1(data [][]int) int {
	result := 0

	for i := 0; i < len(data)/2; i++ {
		result += getScore(findMatchingNumbers(data[i*2], data[i*2+1]))
	}

	return result
}
func getScore(matches int) int {
	if matches == 0 {
		return 0
	}
	return intPow(2, matches-1)
}
func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
func findMatchingNumbers(winningNumbers, myNumbers []int) int {
	result := 0

	for _, number := range myNumbers {
		if slices.Contains(winningNumbers, number) {
			result++
		}
	}

	return result
}

func getData(filename string) [][]int {
	result := [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineData := []int{}

		parts := strings.Split(line, " ")

		for i := 2; i < len(parts); i++ {
			if len(parts[i]) == 0 {
				continue
			}
			if parts[i] == "|" {
				result = append(result, lineData)
				lineData = []int{}
				continue
			}

			value, _ := strconv.ParseInt(parts[i], 10, 32)
			lineData = append(lineData, int(value))
		}

		result = append(result, lineData)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
