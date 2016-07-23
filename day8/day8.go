package main

import (
	"strconv"
	"strings"
)

// UnquatedLengthDiff returns string unquated length diff
func UnquatedLengthDiff(s string) int {
	uq, _ := strconv.Unquote(s)
	return len(s) - len(uq)
}

// QuatedLengthDiff returns string quated length diff
func QuatedLengthDiff(s string) int {
	q := strconv.Quote(s)
	return len(q) - len(s)
}

// DiffCalc calculates a difference according to a given function
func DiffCalc(lines string, fn func(string) int) int {
	var total int
	splitLines := strings.Split(lines, "\n")

	for _, s := range splitLines {
		total += fn(s)
	}

	return total
}
