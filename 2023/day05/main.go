package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	seeds, mappings := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(seeds, mappings))
	fmt.Printf("Part 2: %d\n", part2(seeds, mappings))
}
func part2(seeds []int64, mappings [7][]int64) int64 {
	lowest := int64(math.MaxInt64)

	for i := 0; i < len(seeds)/2; i++ {
		start := seeds[i*2]
		r := seeds[i*2+1]

		location := mapSeedRange(start, r, mappings)

		if location < lowest {
			lowest = location
		}
	}

	return lowest
}
func mapSeedRange(start, r int64, mappings [7][]int64) int64 {
	lowest := int64(math.MaxInt64)

	for i := int64(0); i < r; i++ {
		location := mapSeed(start+i, mappings)

		if location < lowest {
			lowest = location
		}
	}

	return lowest
}
func part1(seeds []int64, mappings [7][]int64) int64 {
	lowest := int64(math.MaxInt64)

	for _, s := range seeds {
		location := mapSeed(s, mappings)

		if location < lowest {
			lowest = location
		}
	}

	return lowest
}
func mapSeed(seed int64, mappings [7][]int64) int64 {
	current := seed

	for i := 0; i < 7; i++ {
		current = mapValue(current, mappings[i])
	}

	return current
}
func mapValue(v int64, mapping []int64) int64 {
	for i := 0; i < len(mapping)/3; i++ {
		destination := mapping[i*3]
		start := mapping[i*3+1]
		r := mapping[i*3+2]

		if v >= start && v < start+r {
			return destination + (v - start)
		}
	}

	return v
}

func getData(filename string) ([]int64, [7][]int64) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	seeds := parseIntegers(line[6:])

	mappings := [7][]int64{}
	currentTarget := -1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map") {
			currentTarget++
			continue
		}

		mappings[currentTarget] = append(mappings[currentTarget], parseIntegers(line)...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seeds, mappings
}
func parseIntegers(line string) []int64 {
	result := []int64{}

	parts := strings.Split(line, " ")
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		value, _ := strconv.ParseInt(p, 10, 64)
		result = append(result, value)
	}

	return result
}
