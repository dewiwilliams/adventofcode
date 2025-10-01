package grid

import "log"

type Grid struct {
	Grid   []int
	Width  int
	Height int
}

func NewFromData(data []string, mapping map[rune]int) Grid {
	if len(data) == 0 {
		log.Fatalf("no data for grid")
	}

	result := Grid{}
	result.Height = len(data)
	result.Width = len(data[0])
	result.Grid = make([]int, result.Width*result.Height)

	for i, line := range data {
		if len(line) != result.Width {
			log.Fatalf("empty line for grid")
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
func (g *Grid) CellCount() int {
	return g.Width * g.Height
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
