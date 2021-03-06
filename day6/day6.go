package main

import (
	"strconv"
	"strings"
)

const stateOn = "on"
const stateOff = "off"
const stateToggle = "toggle"

// Grid system
type Grid struct {
	Map      map[int]map[int]bool
	TurnedOn int
}

// NewGrid system
type NewGrid struct {
	Map      map[int]map[int]int
	TurnedOn int
}

// ChangeLigtState changes the state of the lights according to the coordinates
func (grid *Grid) ChangeLigtState(state string, from string, to string) *Grid {
	fromSplit := strings.Split(from, ",")
	toSplit := strings.Split(to, ",")

	if grid.Map == nil {
		grid.Map = make(map[int]map[int]bool)
	}

	for i, n := strToInt(fromSplit[0]), strToInt(toSplit[0]); i <= n; i++ {
		for j, m := strToInt(fromSplit[1]), strToInt(toSplit[1]); j <= m; j++ {
			if grid.Map[i] == nil {
				grid.Map[i] = make(map[int]bool)
				grid.Map[i][j] = false
			}

			switch state {
			case stateOn:
				if grid.Map[i][j] == false {
					grid.Map[i][j] = true
					grid.TurnedOn++
				}
			case stateOff:
				if grid.Map[i][j] {
					grid.Map[i][j] = false
					grid.TurnedOn--
				}
			case stateToggle:
				if grid.Map[i][j] {
					grid.TurnedOn--
					grid.Map[i][j] = false
				} else {
					grid.TurnedOn++
					grid.Map[i][j] = true
				}
			}
		}
	}

	// fmt.Printf("Action: %s From %s To %s, current: %d\n", state, from, to, grid.TurnedOn)

	return grid
}

// ChangeNewLigtState changes the state of the lights according to the coordinates
func (grid *NewGrid) ChangeNewLigtState(state string, from string, to string) *NewGrid {
	fromSplit := strings.Split(from, ",")
	toSplit := strings.Split(to, ",")

	if grid.Map == nil {
		grid.Map = make(map[int]map[int]int)
	}

	for i, n := strToInt(fromSplit[0]), strToInt(toSplit[0]); i <= n; i++ {
		for j, m := strToInt(fromSplit[1]), strToInt(toSplit[1]); j <= m; j++ {
			if grid.Map[i] == nil {
				grid.Map[i] = make(map[int]int)
				grid.Map[i][j] = 0
			}

			switch state {
			case stateOn:
				grid.TurnedOn++
				grid.Map[i][j]++
			case stateOff:
				if grid.Map[i][j] > 0 {
					grid.Map[i][j]--
					grid.TurnedOn--
				}
			case stateToggle:
				grid.TurnedOn += 2
				grid.Map[i][j] += 2
			}
		}
	}

	// fmt.Printf("Action: %s From %s To %s, current: %d\n", state, from, to, grid.TurnedOn)

	return grid
}

// ParseInstructions Parses a given string instruction
func (grid *Grid) ParseInstructions(instructions string) *Grid {
	rows := strings.Split(instructions, "\n")

	// check each row
	for _, row := range rows {
		row = strings.Trim(row, "	")
		inst := strings.Split(row, " ")
		switch inst[0] {
		case "turn":
			grid.ChangeLigtState(inst[1], inst[2], inst[4])
		case stateToggle:
			grid.ChangeLigtState(inst[0], inst[1], inst[3])
		}
	}
	return grid
}

// ParseNewInstructions Parses a given string instruction
func (grid *NewGrid) ParseNewInstructions(instructions string) *NewGrid {
	rows := strings.Split(instructions, "\n")

	// check each row
	for _, row := range rows {
		row = strings.Trim(row, "	")
		inst := strings.Split(row, " ")
		switch inst[0] {
		case "turn":
			grid.ChangeNewLigtState(inst[1], inst[2], inst[4])
		case stateToggle:
			grid.ChangeNewLigtState(inst[0], inst[1], inst[3])
		}
	}
	return grid
}

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}
