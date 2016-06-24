package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Present is a box with width, height & length properties
type Present struct {
	Width  int
	Height int
	Length int
}

// PresentList is an array of Presents
type PresentList []Present

// GetWrapperDimensions returns how much wrapping paper
// is needed to wrap the present
func (p Present) GetWrapperDimensions() int {
	sides := []int{
		p.Length * p.Width,
		p.Width * p.Height,
		p.Height * p.Length,
	}
	sort.Ints(sides)
	// initialise to the area of the smallest side
	sum := sides[0]
	for _, val := range sides {
		sum += (2 * val)
	}
	return sum
}

// GetRibbonLength returns how much ribbon
// is needed in order wrap the present
func (p Present) GetRibbonLength() int {
	lengths := []int{
		p.Length,
		p.Width,
		p.Height,
	}
	sort.Ints(lengths)
	return (2*lengths[0] + 2*lengths[1]) + (lengths[0] * lengths[1] * lengths[2])
}

// GetTotalWrapperDimensions sums up all of the presents wrapper dimensions
// that were calculated by GetWrapperDimensions
func (ps PresentList) GetTotalWrapperDimensions() int {
	sum := 0
	for _, p := range ps {
		sum += p.GetWrapperDimensions()
	}
	return sum
}

// GetTotalRobbonLength sums up all of the presents ribbon lengths
// that were calculated by GetRibbonLength
func (ps PresentList) GetTotalRibbonLength() int {
	sum := 0
	for _, p := range ps {
		sum += p.GetRibbonLength()
	}
	return sum
}

func main() {
	presents := parseInput(getInput())
	fmt.Printf("Total wrapper dimensions: %d\n", presents.GetTotalWrapperDimensions())
	fmt.Printf("Total ribbon length: %d\n", presents.GetTotalRibbonLength())
}

// parseInput Builds the PresentList from the given input
func parseInput(input string) PresentList {
	presents := strings.Split(input, "\n")
	pl := PresentList{}
	for _, p := range presents {
		dimArray := strings.Split(p, "x")

		// convert strings to integers
		width, _ := strconv.Atoi(dimArray[0])
		height, _ := strconv.Atoi(dimArray[1])
		length, _ := strconv.Atoi(dimArray[2])

		pl = append(pl, Present{
			Width:  width,
			Height: height,
			Length: length,
		})
	}
	return pl
}

// getInput returns the input provided by the game
func getInput() string {
	return `29x13x26`
}
