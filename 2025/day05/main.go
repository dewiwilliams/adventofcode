package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type idRange struct {
	start int64
	end   int64
}

func main() {
	//ranges, ingredients := getData("sample_input.txt")
	ranges, ingredients := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(ranges, ingredients))
	fmt.Printf("Part 2: %d\n", part2(ranges))
}

func part1(ranges []idRange, ingredients []int64) int {
	result := 0

	for _, ingredient := range ingredients {
		if isFresh(ranges, ingredient) {
			result++
		}
	}

	return result
}
func isFresh(ranges []idRange, ingredient int64) bool {
	for _, r := range ranges {
		if ingredient >= r.start && ingredient <= r.end {
			return true
		}
	}

	return false
}

func part2(ranges []idRange) int64 {
	for {
		newRange := mergeRanges(ranges)
		if newRange == nil {
			return countIngredients(ranges)
		}

		ranges = newRange
	}
}
func countIngredients(ranges []idRange) int64 {
	var result int64

	for _, r := range ranges {
		result += r.end - r.start + 1
	}

	return result
}
func mergeRanges(ranges []idRange) []idRange {
	for i, ri := range ranges {
		for j := i + 1; j < len(ranges); j++ {
			rj := ranges[j]

			if !canMerge(ranges[i], ranges[j]) {
				continue
			}

			ranges[i] = idRange{
				start: min(ri.start, rj.start),
				end:   max(ri.end, rj.end),
			}

			return remove(ranges, j)
		}
	}

	return nil
}
func remove(s []idRange, i int) []idRange {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func canMerge(r1, r2 idRange) bool {
	return max(r1.start, r2.start) <= min(r1.end+1, r2.end+1)
}

func getData(filename string) ([]idRange, []int64) {
	lines := util.GetFileLines(filename)

	ranges := []idRange{}
	ingredients := []int64{}

	ingredientStart := -1

	for i, l := range lines {
		if len(l) == 0 {
			ingredientStart = i + 1
			break
		}

		parts := strings.Split(l, "-")

		if len(parts) != 2 {
			log.Fatalf("Multiple parts (%s)", l)
		}

		start, errStart := strconv.ParseInt(parts[0], 10, 64)
		if errStart != nil {
			log.Fatalf("Failed parsing low: %s", parts[0])
		}

		end, errEnd := strconv.ParseInt(parts[1], 10, 64)
		if errEnd != nil {
			log.Fatalf("Failed parsing high: %s", parts[1])
		}

		ranges = append(ranges, idRange{
			start: start,
			end:   end,
		})
	}

	for i := ingredientStart; i < len(lines); i++ {
		ingredient, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			log.Fatalf("Failed parsing ingredient: %s", lines[i])
		}

		ingredients = append(ingredients, ingredient)
	}

	return ranges, ingredients
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
