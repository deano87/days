package main

import "fmt"

// Pointer struct points to the current position in the Scale
// and holds the number of houses found
type Pointer struct {
	X       int
	Y       int
	Counter int // counter for unique houses
}

// Houses represent the infinite 2scale map
type Houses map[int]map[int]int

// HandleDelivery receives an instruction string and returns
// how many houses were visited
func HandleDelivery(h *Houses, p *Pointer, fn func(x int, y int) (int, int)) int {
	// update the pointers
	(*p).X, (*p).Y = fn((*p).X, (*p).Y)

	// check if house already been visited
	if _, ok := (*h)[(*p).X][(*p).Y]; ok {
		(*h)[(*p).X][(*p).Y]++
	} else {
		// add house to the map
		if _, ok = (*h)[(*p).X]; !ok {
			(*h)[(*p).X] = make(map[int]int)
		}
		(*h)[(*p).X][(*p).Y] = 1
		// update the houses counter
		(*p).Counter++
	}
	return (*p).Counter
}

// ParseElfInstructions parses the given string and calls the delivery function
func ParseElfInstructions(s string) int {
	var santaPointer = &Pointer{X: 0, Y: 0, Counter: 1}
	var robotPointer = &Pointer{X: 0, Y: 0, Counter: 0}
	var houses = &Houses{0: {0: 1}} // first house always gets a present

	fmt.Printf("Starting with string: %s\n", s)

	for i := 0; i < len(s); i++ {
		var pointer = santaPointer
		if i%2 == 0 {
			pointer = robotPointer
		}
		switch s[i] {
		case '^':
			HandleDelivery(houses, pointer, moveNorth)
		case '>':
			HandleDelivery(houses, pointer, moveEast)
		case 'v':
			HandleDelivery(houses, pointer, moveSouth)
		case '<':
			HandleDelivery(houses, pointer, moveWest)
		}
	}

	// return how many houses the robot & santa visited
	return santaPointer.Counter + robotPointer.Counter
}

// Move functions

func moveNorth(x int, y int) (int, int) {
	return x, y + 1
}

func moveEast(x int, y int) (int, int) {
	return x + 1, y
}

func moveSouth(x int, y int) (int, int) {
	return x, y - 1
}

func moveWest(x int, y int) (int, int) {
	return x - 1, y
}
