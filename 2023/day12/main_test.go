package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardToInteger(t *testing.T) {
	assert.Equal(t, 1, getCombinations([]int{unknownCell, unknownCell}, []int{2}))
	assert.Equal(t, 2, getCombinations([]int{unknownCell, unknownCell, unknownCell}, []int{2}))
	assert.Equal(t, 1, getCombinations([]int{unknownCell, unknownCell, unknownCell}, []int{1, 1}))
	assert.Equal(t, 3, getCombinations([]int{unknownCell, unknownCell, unknownCell}, []int{1}))
	assert.Equal(t, 0, getCombinations([]int{unknownCell, emptyCell, unknownCell}, []int{2}))
	assert.Equal(t, 2, getCombinations([]int{unknownCell, unknownCell, filledCell, unknownCell}, []int{2}))

	mapping := getRecordMapping()
	assert.Equal(t, 4, getCombinations(mapStringToArray(".??..??...?##.", mapping), []int{1, 1, 3}))
}
