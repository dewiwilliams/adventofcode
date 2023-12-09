package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardToInteger(t *testing.T) {
	assert.Equal(t, 2, cardToInteger('2'), "2")
	assert.Equal(t, 9, cardToInteger('9'), "9")
	assert.Equal(t, 10, cardToInteger('T'), "T")
	assert.Equal(t, 11, cardToInteger('J'), "J")
	assert.Equal(t, 12, cardToInteger('Q'), "Q")
	assert.Equal(t, 13, cardToInteger('K'), "K")
	assert.Equal(t, 14, cardToInteger('A'), "A")
}
func TestGetHandType(t *testing.T) {
	assert.Equal(t, FIVE_OF_A_KIND, getHandType([]rune("AAAAA")), "Five of a kind")
	assert.Equal(t, FOUR_OF_A_KIND, getHandType([]rune("AA8AA")), "Four of a kind")
	assert.Equal(t, FULL_HOUSE, getHandType([]rune("23332")), "Full house")
	assert.Equal(t, THREE_OF_A_KIND, getHandType([]rune("TTT98")), "Three of a kind")
	assert.Equal(t, TWO_PAIR, getHandType([]rune("23432")), "Two pair")
	assert.Equal(t, ONE_PAIR, getHandType([]rune("A23A4")), "One pair")
	assert.Equal(t, HIGH_CARD, getHandType([]rune("23456")), "High card")
}
func TestGetHandCardScore(t *testing.T) {
	assert.Equal(t, 0xDD677, getHandCardScore([]rune("KK677")), "KK677")
}
