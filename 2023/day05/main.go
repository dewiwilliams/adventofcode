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
func splitRange(seed []int64, mapping []int64) []int64 {
	seedStart := seed[0]
	seedLength := seed[1]
	seedEnd := seedStart + seedLength
	mappingStart := mapping[0]
	mappingLength := mapping[1]
	mappingEnd := mappingStart + mappingLength

	if seedEnd < mappingStart {
		return []int64{seedStart, seedLength}
	}
	if seedStart > mappingEnd {
		return []int64{seedStart, seedLength}
	}
	if seedStart < mappingStart && seedEnd > mappingStart {
		firstSectionLength := mappingStart - seedStart

		if seedEnd > mappingEnd {
			return []int64{
				seedStart, firstSectionLength,
				mappingStart, mappingLength,
				mappingEnd, seedLength - firstSectionLength - mappingLength,
			}
		} else {
			return []int64{
				seedStart, firstSectionLength,
				mappingStart, seedLength - firstSectionLength,
			}
		}
	}
	if seedStart >= mappingStart && seedStart < mappingEnd && seedEnd > mappingEnd {
		return []int64{
			seedStart, mappingEnd - seedStart,
			mappingEnd, seedLength - (mappingEnd - seedStart),
		}
	}

	return []int64{seedStart, seedLength}
}
func splitRanges(seed []int64, mappings []int64) []int64 {
	for i := 0; i < len(mappings)/3; i++ {
		mapping := []int64{mappings[i*3+1], mappings[i*3+2]}
		splitSeeds := splitRange(seed, mapping)
		if len(splitSeeds) == 2 {
			continue
		}

		result := []int64{}
		for j := 0; j < len(splitSeeds)/2; j++ {
			result = append(result, splitRanges([]int64{splitSeeds[j*2], splitSeeds[j*2+1]}, mappings)...)
		}

		return result
	}

	return seed
}
func splitSeeds(seeds []int64, mappings []int64) []int64 {
	result := []int64{}

	for i := 0; i < len(seeds)/2; i++ {
		seed := []int64{seeds[i*2], seeds[i*2+1]}
		result = append(result, splitRanges(seed, mappings)...)
	}

	return result
}
func mapSeeds(seeds []int64, mappings []int64) []int64 {
	splitSeeds := splitSeeds(seeds, mappings)
	result := []int64{}

	for i := 0; i < len(splitSeeds)/2; i++ {
		result = append(result, mapValue(splitSeeds[i*2], mappings), splitSeeds[i*2+1])
	}

	return result
}
func part2(seeds []int64, mappings [7][]int64) int64 {
	currentSeeds := seeds
	for i := 0; i < 7; i++ {
		currentSeeds = mapSeeds(currentSeeds, mappings[i])
	}

	lowest := int64(math.MaxInt64)
	for i := 0; i < len(currentSeeds)/2; i++ {
		location := currentSeeds[i*2]

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
		length := mapping[i*3+2]

		if v >= start && v < start+length {
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
