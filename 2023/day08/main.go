package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

func main() {
	directions, nodes := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(directions, nodes))
	fmt.Printf("Part 2: %d\n", part2(directions, nodes))
}
func part2(directions []rune, nodeData []string) int64 {
	nodes := buildNodesMapData(nodeData)
	startingLocations := getStartingLocations(nodeData)

	cycleOffsets := []int64{}
	cycleLengths := []int64{}

	for _, location := range startingLocations {
		offset, length := getCycleLength(location, directions, nodes)

		if offset != length {
			// This is an observation from the data.
			panic("cycle offset and length are not equal!")
		}
		cycleOffsets = append(cycleOffsets, int64(offset))
		cycleLengths = append(cycleLengths, int64(length))
	}

	return lcm(cycleLengths[0], cycleLengths[1], cycleLengths[2:]...)
}
func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int64, integers ...int64) int64 {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
func getLocationAfterSteps(startingLocation string, directions []rune, nodes map[string]node, steps int64) string {
	currentLocation := startingLocation
	directionLength := int64(len(directions))

	for i := int64(0); i < steps; i++ {
		directionIndex := i % directionLength

		currentLocation = advance(nodes, currentLocation, directions[directionIndex])
	}

	return currentLocation
}
func getCycleLength(startingLocation string, directions []rune, nodes map[string]node) (int, int) {
	offset := 0
	lastCycleIndex := 0
	currentLocation := startingLocation
	directionLength := len(directions)

	for step := 0; ; step++ {
		directionIndex := step % directionLength

		currentLocation = advance(nodes, currentLocation, directions[directionIndex])
		if currentLocation[2] == 'Z' {
			if offset == 0 {
				offset = step + 1
				lastCycleIndex = step
			} else {
				return offset, step - lastCycleIndex
			}
		}
	}
}
func getStartingLocations(nodeData []string) []string {
	result := []string{}

	for i := 0; i < len(nodeData)/3; i++ {
		if nodeData[i*3][2] == 'A' {
			result = append(result, nodeData[i*3])
		}
	}

	return result
}
func getEndNodeCount(nodes []string) int {
	result := 0

	for _, node := range nodes {
		if node[2] == 'Z' {
			result++
		}
	}

	return result
}
func areAllEndNodes(nodes []string) bool {
	return getEndNodeCount(nodes) == len(nodes)
}
func part1(directions []rune, nodeData []string) int {

	nodes := buildNodesMapData(nodeData)
	currentLocation := "AAA"
	directionLength := len(directions)

	for step := 0; ; step++ {
		directionIndex := step % directionLength
		currentLocation = advance(nodes, currentLocation, directions[directionIndex])

		if currentLocation == "ZZZ" {
			return step + 1
		}
	}
}
func advance(nodes map[string]node, currentLocation string, direction rune) string {
	if direction == 'L' {
		return nodes[currentLocation].left
	} else if direction == 'R' {
		return nodes[currentLocation].right
	} else {
		panic("direction is not left or right")
	}
}
func buildNodesMapData(nodeData []string) map[string]node {
	nodes := make(map[string]node)
	for i := 0; i < len(nodeData)/3; i++ {
		nodes[nodeData[i*3]] = node{
			left:  nodeData[i*3+1],
			right: nodeData[i*3+2],
		}
	}

	return nodes
}
func parseData(filename string) ([]rune, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := []string{}

	scanner.Scan()
	directions := strings.TrimSpace(scanner.Text())

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		runes := []rune(line)

		nodes = append(nodes,
			string([]rune{runes[0], runes[1], runes[2]}),
			string([]rune{runes[7], runes[8], runes[9]}),
			string([]rune{runes[12], runes[13], runes[14]}),
		)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return []rune(directions), nodes
}
