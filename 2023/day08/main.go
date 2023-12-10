package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	directions, nodes := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(directions, nodes))
}
func part1(directions []rune, nodeData []string) int {
	type node struct {
		left  string
		right string
	}
	nodes := make(map[string]node)
	for i := 0; i < len(nodeData)/3; i++ {
		nodes[nodeData[i*3]] = node{
			left:  nodeData[i*3+1],
			right: nodeData[i*3+2],
		}
	}

	currentLocation := "AAA"
	directionLength := len(directions)

	for step := 0; ; step++ {
		directionIndex := step % directionLength
		if directions[directionIndex] == 'L' {
			currentLocation = nodes[currentLocation].left
		} else if directions[directionIndex] == 'R' {
			currentLocation = nodes[currentLocation].right
		} else {
			panic("direction is not left or right")
		}

		if currentLocation == "ZZZ" {
			return step + 1
		}
	}
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
