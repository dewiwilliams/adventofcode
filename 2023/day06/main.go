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
	times, distances := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(times, distances))
	fmt.Printf("Part 2: %d\n", part2(times, distances))
}
func part2(times, distances []string) int64 {
	time, _ := strconv.ParseInt(strings.Join(times, ""), 10, 64)
	distance, _ := strconv.ParseInt(strings.Join(distances, ""), 10, 64)

	return getDifferentWinningTimes(int64(time), int64(distance))
}
func part1(stringTimes, stringDistances []string) int64 {
	times := parseArray(stringTimes)
	distances := parseArray(stringDistances)

	result := int64(1)

	for i := 0; i < len(times); i++ {
		result *= getDifferentWinningTimes(times[i], distances[i])
	}

	return result
}
func getDifferentWinningTimes(time, winningDistance int64) int64 {
	result := int64(0)

	for t := int64(1); t < time; t++ {
		distance := getDistance(t, time)
		if distance > winningDistance {
			result++
		}
	}

	return result
}
func getDistance(holdTime, totalTime int64) int64 {
	return holdTime * (totalTime - holdTime)
}
func getData(filename string) ([]string, []string) {
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

		parts := strings.Fields(line[11:])
		result = append(result, parts...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	midpoint := len(result) / 2
	return result[:midpoint], result[midpoint:]
}
func parseArray(data []string) []int64 {
	result := []int64{}

	for _, d := range data {
		value, _ := strconv.ParseInt(d, 10, 64)
		result = append(result, value)
	}

	return result
}
