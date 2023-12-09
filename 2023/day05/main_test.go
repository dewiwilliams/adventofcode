package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitRange(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int64{0, 10}, splitRange([]int64{0, 10}, []int64{10, 5}), "No overlap lower")
	assert.Equal([]int64{15, 9}, splitRange([]int64{15, 9}, []int64{10, 5}), "No overlap higher")
	assert.Equal([]int64{0, 10, 10, 10}, splitRange([]int64{0, 20}, []int64{10, 10}), "Overlap lower")
	assert.Equal([]int64{10, 10, 20, 10}, splitRange([]int64{10, 20}, []int64{10, 10}), "Overlap higher")
	assert.Equal([]int64{0, 10, 10, 10, 20, 10}, splitRange([]int64{0, 30}, []int64{10, 10}), "Overlap both ends")
}
func TestSplitRanges(t *testing.T) {
	assert := assert.New(t)

	seed := []int64{0, 30}
	mappings := []int64{
		0, 5, 5,
		0, 11, 9,
		0, 20, 5,
	}
	expectedResult := []int64{0, 5, 5, 5, 10, 1, 11, 9, 20, 5, 25, 5}

	assert.Equal(expectedResult, splitRanges(seed, mappings), "Split ranges")
}
