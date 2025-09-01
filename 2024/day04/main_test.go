package main

import (
	"slices"
	"testing"
)

func TestGetStrides(t *testing.T) {
	width := 10
	height := 20

	data := []struct {
		x               int
		y               int
		expectedStrides []int
	}{
		{
			x:               2,
			y:               2,
			expectedStrides: []int{1, width, width + 1},
		},
		{
			x:               3,
			y:               3,
			expectedStrides: []int{1, width, width + 1, -1, width - 1, -width, -width - 1, -width + 1},
		},
		{
			x:               6,
			y:               16,
			expectedStrides: []int{1, width, width + 1, -1, width - 1, -width, -width - 1, -width + 1},
		},
		{
			x:               7,
			y:               17,
			expectedStrides: []int{-1, -width, -width - 1},
		},
	}

	for _, d := range data {
		strides := getStrides(d.x, d.y, width, height)
		slices.Sort(strides)
		slices.Sort(d.expectedStrides)

		if !slices.Equal(strides, d.expectedStrides) {
			t.Errorf("Mismatched strides (%d, %d): %v vs %v\n", d.x, d.y, strides, d.expectedStrides)
		}
	}
}
