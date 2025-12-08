package main

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
	z int
}
type linkedBoxes struct {
	index1          int
	index2          int
	squaredDistance int
}
type circuit struct {
	boxIndeces []int
}

func main() {
	//filename := "sample_input.txt"
	filename := "input.txt"
	data := getData(filename)

	fmt.Printf("Part 1: %d\n", part1(data, 1000))
	//fmt.Printf("Part 2: %d\n", part2(filename))
}

func part1(coordinates []coordinate, iterations int) int {
	linkedBoxes := buildLinkedBoxes(coordinates)

	sort.Slice(linkedBoxes, func(i, j int) bool {
		return linkedBoxes[i].squaredDistance < linkedBoxes[j].squaredDistance
	})

	circuits := []circuit{}

	for i := range iterations {
		circuits = append(circuits, circuit{
			boxIndeces: []int{
				linkedBoxes[i].index1,
				linkedBoxes[i].index2,
			},
		})
	}

	mergeCircuits(circuits)

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].boxIndeces) > len(circuits[j].boxIndeces)
	})

	return len(circuits[0].boxIndeces) * len(circuits[1].boxIndeces) * len(circuits[2].boxIndeces)
}
func part2(coordinates []coordinate) int {
	linkedBoxes := buildLinkedBoxes(coordinates)

	sort.Slice(linkedBoxes, func(i, j int) bool {
		return linkedBoxes[i].squaredDistance < linkedBoxes[j].squaredDistance
	})

	circuits := []circuit{}

	for i := range linkedBoxes {
		circuits = append(circuits, circuit{
			boxIndeces: []int{
				linkedBoxes[i].index1,
				linkedBoxes[i].index2,
			},
		})

		mergeCircuits(circuits)
	}

	return 0
}
func mergeCircuits(circuits []circuit) []circuit {
	limit := len(circuits)

	for i := 0; i < limit; i++ {
		for j := i + 1; j < limit; j++ {
			if hasIntersection(circuits[i].boxIndeces, circuits[j].boxIndeces) {
				merged := append(circuits[j].boxIndeces, circuits[i].boxIndeces...)
				slices.Sort(merged)
				circuits[j].boxIndeces = slices.Compact(merged)
				circuits[i].boxIndeces = []int{}
			}
		}
	}

	return circuits
}
func hasIntersection(a, b []int) bool {
	for _, v1 := range a {
		for _, v2 := range b {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}
func buildLinkedBoxes(coordinates []coordinate) []linkedBoxes {
	result := []linkedBoxes{}

	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			xDiff := coordinates[i].x - coordinates[j].x
			yDiff := coordinates[i].y - coordinates[j].y
			zDiff := coordinates[i].z - coordinates[j].z

			result = append(result, linkedBoxes{
				index1:          i,
				index2:          j,
				squaredDistance: xDiff*xDiff + yDiff*yDiff + zDiff*zDiff,
			})
		}
	}

	return result
}

func getData(filename string) []coordinate {
	lines := util.GetFileLines(filename)

	result := []coordinate{}

	for _, l := range lines {
		parts := strings.Split(l, ",")

		x, _ := strconv.ParseInt(parts[0], 10, 32)
		y, _ := strconv.ParseInt(parts[1], 10, 32)
		z, _ := strconv.ParseInt(parts[2], 10, 32)

		result = append(result, coordinate{
			x: int(x),
			y: int(y),
			z: int(z),
		})
	}

	return result
}
