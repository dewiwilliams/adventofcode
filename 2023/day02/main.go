package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	red   int
	green int
	blue  int
}

type game struct {
	number int
	hands  []hand
}

func main() {
	data := getData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(data))
	fmt.Printf("Part 2: %d\n", part2(data))
}

func part1(games []game) int {
	result := 0

	for _, g := range games {
		if isGameValid(g) {
			result += g.number
		}
	}

	return result
}
func isGameValid(game game) bool {
	for _, h := range game.hands {
		if !isHandValid(h) {
			return false
		}
	}

	return true
}
func isHandValid(hand hand) bool {
	// 12 red cubes, 13 green cubes, and 14 blue cubes

	return hand.red <= 12 &&
		hand.green <= 13 &&
		hand.blue <= 14
}
func part2(games []game) int {
	result := 0

	for _, g := range games {
		result += getHandPower(getMinimumHand(g.hands))
	}

	return result
}
func getHandPower(hand hand) int {
	return hand.blue * hand.green * hand.red
}
func getMinimumHand(hands []hand) hand {
	result := hands[0]

	for i := 1; i < len(hands); i++ {
		result.red = max(result.red, hands[i].red)
		result.green = max(result.green, hands[i].green)
		result.blue = max(result.blue, hands[i].blue)
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func getData(filename string) []game {
	result := []game{}

	file, err := os.Open(filename)
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

		result = append(result, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
func parseLine(line string) game {
	result := game{}

	parts := strings.FieldsFunc(line, split)

	gameNumber, _ := strconv.ParseInt(parts[0][5:], 10, 32)
	result.number = int(gameNumber)

	for i := 1; i < len(parts); i++ {
		result.hands = append(result.hands, parseGame(strings.TrimSpace(parts[i])))
	}

	return result
}
func split(r rune) bool {
	return r == ':' || r == ';'
}
func parseGame(game string) hand {
	result := hand{}

	parts := strings.Split(game, ",")

	for _, p := range parts {
		s := strings.Split(strings.TrimSpace(p), " ")
		q, _ := strconv.ParseInt(s[0], 10, 32)

		if s[1] == "green" {
			result.green = int(q)
		} else if s[1] == "blue" {
			result.blue = int(q)
		} else if s[1] == "red" {
			result.red = int(q)
		}
	}

	return result
}
