package main

import (
	"fmt"
	"strings"
)

// CountGoodStrings checks how many given strings are good
func CountGoodStrings(s string, funcs ...func(string) bool) int {
	rows := strings.Split(s, "\n")
	count := 0
	// check each row
	for _, row := range rows {
		flag := true
		row = strings.Trim(row, " ")
		fmt.Printf("handling: %s ...", row)
		// check all provided functions
		for _, f := range funcs {
			if !f(row) {
				flag = false
				fmt.Printf("bad\n")
			}
		}
		if flag {
			fmt.Printf("good\n")
			count++
		}
	}
	return count
}

// Good string functions

// ThreeVowels - Good string has at least 3 vowels
func ThreeVowels(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

// DoubleLetter - string is good if it has a least one double letter
func DoubleLetter(s string) bool {
	lastLetter := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == lastLetter {
			return true
		}
		lastLetter = s[i]
	}
	return false
}

// NoBadCombinations - string is good if it doesn't contain bad strings
func NoBadCombinations(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		switch s[i:(i + 2)] {
		case "ab":
			fallthrough
		case "cd":
			fallthrough
		case "pq":
			fallthrough
		case "xy":
			return false
		}
	}
	return true
}

// DoubleChars - checks if two chars appear twice in the string
func DoubleChars(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		check := s[i : i+2]
		for j := i + 2; j < len(s)-1; j++ {
			if check == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

// LetterRepeat - checks if a letter repeats with one letter between
func LetterRepeat(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
