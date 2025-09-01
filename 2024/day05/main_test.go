package main

import (
	"slices"
	"testing"
)

func TestIsOrderingCorrect(t *testing.T) {
	ordering := make(map[int][]int)
	ordering[29] = []int{13}
	ordering[47] = []int{53, 13, 61, 29}
	ordering[53] = []int{29, 13}
	ordering[61] = []int{13, 53, 29}
	ordering[75] = []int{29, 53, 47, 61, 13}
	ordering[97] = []int{13, 61, 47, 29, 53, 75}

	reverseOrdering := make(map[int][]int)
	reverseOrdering[13] = []int{53}
	reverseOrdering[29] = []int{13, 47}
	reverseOrdering[47] = []int{53, 13, 75}
	reverseOrdering[53] = []int{29, 97}
	reverseOrdering[61] = []int{13, 53, 29, 75}
	reverseOrdering[75] = []int{29, 53, 47, 61}

	correctOrderPrints := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}
	for _, v := range correctOrderPrints {
		if !isOrderingCorrect(ordering, v) {
			t.Fatalf("Ordering not correct!")
		}
	}

	incorrectOrderPrints := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	for _, v := range incorrectOrderPrints {
		if isOrderingCorrect(ordering, v) {
			t.Fatalf("Ordering is correct!")
		}
	}
}
func TestCorrectOrdering(t *testing.T) {
	ordering := make(map[int][]int)
	ordering[29] = []int{13}
	ordering[47] = []int{53, 13, 61, 29}
	ordering[53] = []int{29, 13}
	ordering[61] = []int{13, 53, 29}
	ordering[75] = []int{29, 53, 47, 61, 13}
	ordering[97] = []int{13, 61, 47, 29, 53, 75}

	data := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	expectedData := [][]int{
		{97, 75, 47, 61, 53},
		{61, 29, 13},
		{97, 75, 47, 29, 13},
	}
	for i, v := range data {
		correctOrder := correctOrdering(ordering, v)
		if slices.Equal(correctOrder, expectedData[i]) {
			t.Fatalf("Failed to correct ordering!")
		}
	}
}
