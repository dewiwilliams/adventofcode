package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const FIVE_OF_A_KIND = 1
const FOUR_OF_A_KIND = 2
const FULL_HOUSE = 3
const THREE_OF_A_KIND = 4
const TWO_PAIR = 5
const ONE_PAIR = 6
const HIGH_CARD = 7

func main() {
	hands, bets := parseData("input.txt")

	fmt.Printf("Part 1: %d\n", part1(hands, bets))
}
func part1(hands []rune, bets []int) int {
	type pair struct {
		hand  string
		score int
		bet   int
	}
	pairs := []pair{}

	for i := 0; i < len(bets); i++ {

		hand := hands[i*5 : i*5+5]
		pairs = append(pairs, pair{
			hand:  string(hand),
			score: scoreHand(hand),
			bet:   bets[i],
		})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].score < pairs[j].score
	})

	result := 0
	for i, p := range pairs {
		result += (i + 1) * p.bet
	}

	return result
}
func scoreHand(hand []rune) int {
	return getHandTypeBaseScore(getHandType(hand)) | getHandCardScore(hand)
}
func getHandCardScore(hand []rune) int {
	return cardToInteger(hand[0])<<(4*4) +
		cardToInteger(hand[1])<<(3*4) +
		cardToInteger(hand[2])<<(2*4) +
		cardToInteger(hand[3])<<(1*4) +
		cardToInteger(hand[4])<<(0*4)
}
func getHandTypeBaseScore(handType int) int {
	if handType == FIVE_OF_A_KIND {
		return 0x600000
	} else if handType == FOUR_OF_A_KIND {
		return 0x500000
	} else if handType == FULL_HOUSE {
		return 0x400000
	} else if handType == THREE_OF_A_KIND {
		return 0x300000
	} else if handType == TWO_PAIR {
		return 0x200000
	} else if handType == ONE_PAIR {
		return 0x100000
	}

	return 0
}
func getHandType(cards []rune) int {
	counts := getCounts(cards)

	if counts[0] == 5 {
		return FIVE_OF_A_KIND
	} else if counts[0] == 4 {
		return FOUR_OF_A_KIND
	} else if counts[0] == 3 && counts[1] == 2 {
		return FULL_HOUSE
	} else if counts[0] == 3 && counts[1] == 1 {
		return THREE_OF_A_KIND
	} else if counts[0] == 2 && counts[1] == 2 {
		return TWO_PAIR
	} else if counts[0] == 2 && counts[1] == 1 {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}
}
func getCounts(cards []rune) []int {
	counts := make([]int, 15)

	for _, c := range cards {
		counts[cardToInteger(c)]++
	}

	slices.Sort(counts)
	slices.Reverse(counts)

	return counts
}
func cardToInteger(r rune) int {
	if r < 58 {
		return int(r) - 48
	} else if r == 'T' {
		return 0xA
	} else if r == 'J' {
		return 0xB
	} else if r == 'Q' {
		return 0xC
	} else if r == 'K' {
		return 0xD
	} else if r == 'A' {
		return 0xE
	}

	panic("Card out of range")
}
func parseData(filename string) ([]rune, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := []rune{}
	bets := []int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)

		cards = append(cards, []rune(parts[0])...)

		bet, _ := strconv.ParseInt(parts[1], 10, 32)
		bets = append(bets, int(bet))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cards, bets
}
