package main

import "testing"

var testPart1 = struct {
	strings string
	goods   int
}{

	strings: `ugknbfddgicrmopn
            aaa
            jchzalrnumimnmhp
            haegwjzuvuyypxyu
            dvszwmarrgswjxmb`,
	goods: 2,
}

var testPart2 = struct {
	strings string
	goods   int
}{

	strings: `qjhvhtzxzqqjkmpb
            xxyxx
            uurcxstgmygtbstg
            ieodomkazucvgmuy`,
	goods: 2,
}

func TestStringsPart1(t *testing.T) {
	goodStringsFound := CountGoodStrings(testPart1.strings, ThreeVowels, DoubleLetter, NoBadCombinations)
	if testPart1.goods != goodStringsFound {
		t.Errorf("Number of good strings found (%d) doesn't match what was expected (%d)", goodStringsFound, testPart1.goods)
	}
}

func TestStringsPart2(t *testing.T) {
	goodStringsFound := CountGoodStrings(testPart2.strings, DoubleChars, LetterRepeat)
	if testPart2.goods != goodStringsFound {
		t.Errorf("Number of good strings found (%d) doesn't match what was expected (%d)", goodStringsFound, testPart2.goods)
	}
}
