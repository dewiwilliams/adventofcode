package grid

import (
	"fmt"
	"log"
)

type Grid struct {
	Grid           []int
	mapping        map[rune]int
	inverseMapping map[int]rune
	Width          int
	Height         int
}

func NewFromData(data []string, mapping map[rune]int) Grid {
	if len(data) == 0 {
		log.Fatalf("no data for grid")
	}

	result := Grid{}
	result.Height = len(data)
	result.Width = len(data[0])
	result.Grid = make([]int, result.Width*result.Height)
	result.mapping = mapping

	result.inverseMapping = make(map[int]rune)
	for k, v := range mapping {
		result.inverseMapping[v] = k
	}

	for i, line := range data {
		if len(line) != result.Width {
			log.Fatalf("empty line for grid(%d): %d vs %d", i, len(line), result.Width)
		}

		for j, r := range line {
			if value, ok := mapping[r]; ok {
				result.Grid[i*result.Width+j] = value
			} else {
				log.Fatalf("value not in mapping")
			}
		}
	}

	return result
}

func (g *Grid) GetNeighbours(cell int) []int {
	result := []int{}

	x := cell % g.Width
	y := cell / g.Width

	if x > 0 {
		result = append(result, (x-1)+y*g.Width)
	}
	if x < g.Width-1 {
		result = append(result, (x+1)+y*g.Width)
	}
	if y > 0 {
		result = append(result, x+(y-1)*g.Width)
	}
	if y < g.Height-1 {
		result = append(result, x+(y+1)*g.Width)
	}

	return result
}
func (g *Grid) GetDiagonalNeighbours(cell int) []int {
	result := []int{}

	x := cell % g.Width
	y := cell / g.Width

	if x > 0 {
		result = append(result, (x-1)+y*g.Width)
	}
	if x < g.Width-1 {
		result = append(result, (x+1)+y*g.Width)
	}
	if y > 0 {
		result = append(result, x+(y-1)*g.Width)
	}
	if y < g.Height-1 {
		result = append(result, x+(y+1)*g.Width)
	}

	if x > 0 && y > 0 {
		result = append(result, (x-1)+(y-1)*g.Width)
	}
	if x < g.Width-1 && y > 0 {
		result = append(result, (x+1)+(y-1)*g.Width)
	}
	if x > 0 && y < g.Height-1 {
		result = append(result, (x-1)+(y+1)*g.Width)
	}
	if x < g.Width-1 && y < g.Height-1 {
		result = append(result, (x+1)+(y+1)*g.Width)
	}

	return result
}
func (g *Grid) GetCellsWithValue(value int) []int {
	result := []int{}

	for i, v := range g.Grid {
		if v == value {
			result = append(result, i)
		}
	}

	return result
}
func (g *Grid) GetCellsInRow(row int) []int {
	result := []int{}

	for x := range g.Width {
		result = append(result, x+row*g.Width)
	}

	return result
}
func (g *Grid) GetCellsWithValueInRow(row, value int) []int {
	result := []int{}

	for x := range g.Width {
		cell := x + row*g.Width
		v := g.Grid[cell]

		if v == value {
			result = append(result, cell)
		}
	}

	return result
}
func (g Grid) String() string {
	result := fmt.Sprintf("%dx%d\n", g.Width, g.Height)

	for y := range g.Height {
		for x := range g.Width {
			value := g.Grid[x+y*g.Width]
			result += string(g.inverseMapping[value])
		}
		result += "\n"
	}

	return result
}
func (g *Grid) CellCount() int {
	return g.Width * g.Height
}
func (g *Grid) GetContiguousAreas(start int) []int {
	/*result := make(map[int]bool)

	toProcess := []int{start}
	for len(toProcess) > 0 {
		cell := toProcess
	}

	return result*/

	return nil
}
func GetNumericMapping(target map[rune]int) {
	target['0'] = 0
	target['1'] = 1
	target['2'] = 2
	target['3'] = 3
	target['4'] = 4
	target['5'] = 5
	target['6'] = 6
	target['7'] = 7
	target['8'] = 8
	target['9'] = 9
}
func GetUppercaseAlphaMapping(target map[rune]int) {
	target['A'] = 10
	target['B'] = 11
	target['C'] = 12
	target['D'] = 13
	target['E'] = 14
	target['F'] = 15
	target['G'] = 16
	target['H'] = 17
	target['I'] = 18
	target['J'] = 19
	target['K'] = 20
	target['L'] = 21
	target['M'] = 22
	target['N'] = 23
	target['O'] = 24
	target['P'] = 25
	target['Q'] = 26
	target['R'] = 27
	target['S'] = 28
	target['T'] = 29
	target['U'] = 30
	target['V'] = 31
	target['W'] = 32
	target['X'] = 33
	target['Y'] = 34
	target['Z'] = 35
}
