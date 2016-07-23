package main

import "testing"

var test = struct {
	strings   string
	expected1 int
	expected2 int
}{
	strings: `"sjdivfriyaaqa\xd2v\"k\"mpcu\"yyu\"en"
					"vcqc"
					"zbcwgmbpijcxu\"yins\"sfxn"
					"yumngprx"
					"bbdj"
					"czbggabkzo\"wsnw\"voklp\"s"
					"acwt"`,
	expected1: 122,
	expected2: 77,
}

func TestPart1(t *testing.T) {
	val := DiffCalc(test.strings, UnquatedLengthDiff)
	if test.expected1 != val {
		t.Errorf("Received val (%d) doesn't match what was expected (%d)", val, test.expected1)
	}
}

func TestPart2(t *testing.T) {
	val := DiffCalc(test.strings, QuatedLengthDiff)
	if test.expected2 != val {
		t.Errorf("Received val (%d) doesn't match what was expected (%d)", val, test.expected2)
	}
}
