package main

import "testing"

var testPart1 = struct {
	strings  string
	expected int
}{

	strings: `turn on 0,0 through 999,999
	toggle 0,0 through 999,0
	turn off 499,499 through 500,500`,
	expected: 998996,
}

func TestStringsPart1(t *testing.T) {
	grid := Grid{}
	on := grid.ParseInstructions(testPart1.strings).TurnedOn
	if testPart1.expected != on {
		t.Errorf("Number of lights on (%d) doesn't match what was expected (%d)", on, testPart1.expected)
	}
}
