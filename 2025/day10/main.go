package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

type machine struct {
	targetState         int
	buttons             []int
	joltageRequirements []int
}
type machineState struct {
	state         int
	buttonPresses int
}

func main() {
	//data := getData("sample_input.txt")
	data := getData("input.txt")

	//fmt.Printf("Got data: %v\n", data)

	fmt.Printf("Part 1: %d\n", part1(data))
}

func part1(machines []machine) int {
	result := 0

	for _, m := range machines {
		result += part1Machine(m)
	}

	return result
}
func part1Machine(m machine) int {
	state := []machineState{
		{},
	}
	index := 0

	for {
		currentState := state[index]

		for _, b := range m.buttons {
			newState := (currentState.state ^ b)
			if newState == m.targetState {
				return currentState.buttonPresses + 1
			}

			state = append(state, machineState{
				state:         newState,
				buttonPresses: currentState.buttonPresses + 1,
			})
		}

		index++
	}
}

func getData(filename string) []machine {
	lines := util.GetFileLines(filename)

	result := []machine{}

	for _, line := range lines {
		parts := strings.Fields(line)

		m := machine{}

		stateCount := len(parts[0]) - 2
		for i := 1; i < len(parts[0])-1; i++ {
			if parts[0][i] == '.' {
				continue
			}

			m.targetState |= 1 << (stateCount - i)
		}

		for i := 1; i < len(parts)-1; i++ {
			buttonList := parts[i]

			buttons := util.ParseIntegerArray(buttonList[1:len(buttonList)-1], ",")
			mask := 0
			for _, b := range buttons {
				mask |= 1 << (stateCount - b - 1)
			}

			m.buttons = append(m.buttons, mask)
		}

		j := parts[len(parts)-1]
		m.joltageRequirements = util.ParseIntegerArray(j[1:len(j)-1], ",")

		result = append(result, m)
	}

	return result
}
