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
	count int64
}

func main() {
	//data := getData("sample_input.txt")
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part1(data []idRange) int64 {
	var result int64

	for _, r := range data {
		for i := r.start; i <= r.end; i++ {
			if part1IsValid(i) {
				result += i
			}
		}
	}

	return result
}
func part1IsValid(value int64) bool {
	asString := strconv.FormatInt(value, 10)
	if len(asString)%2 == 1 {
		return false
	}

	halfLength := len(asString) / 2
	for i := 0; i < halfLength; i++ {
		if asString[i] != asString[halfLength+i] {
			return false
		}
	}

	return true
}

func part2(data []idRange) int64 {
	var result int64

	for _, r := range data {
		for i := r.start; i <= r.end; i++ {
			if part2IsValid(i) {
				result += i
			}
		}
	}

	return result
}
func part2IsValid(value int64) bool {
	asString := strconv.FormatInt(value, 10)

	for i := 1; i <= len(asString)/2; i++ {
		if len(asString)%i != 0 {
			continue
		}

		if isRepeated(asString, i) {
			return true
		}
	}

	return false
}
func isRepeated(s string, length int) bool {
	limit := len(s) / length

	for j := 1; j < limit; j++ {
		base := j * length

		for i := range length {
			if s[i] != s[base+i] {
				return false
			}
		}
	}

	return true
}

func getData(filename string) []idRange {
	lines := util.GetFileLines(filename)

	if len(lines) != 1 {
		log.Fatalf("Multiple lines (%d)", len(lines))
	}

	ranges := strings.Split(lines[0], ",")
	result := []idRange{}

	var rangeTotal int64

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		if len(parts) != 2 {
			log.Fatalf("Multiple parts (%s)", r)
		}

		start, errStart := strconv.ParseInt(parts[0], 10, 64)
		if errStart != nil {
			log.Fatalf("Failed parsing low: %s", parts[0])
		}

		end, errEnd := strconv.ParseInt(parts[1], 10, 64)
		if errEnd != nil {
			log.Fatalf("Failed parsing high: %s", parts[1])
		}

		result = append(result, idRange{
			start: start,
			end:   end,
			count: end - start + 1,
		})
		rangeTotal += end - start + 1
	}

	return result
}
