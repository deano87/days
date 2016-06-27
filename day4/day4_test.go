package main

import "testing"

var testPart1 = []struct {
	key    string
	result int32
}{
	{
		key:    "abcdef",
		result: 609043,
	},
	{
		key:    "pqrstuv",
		result: 1048970,
	},
	{
		key:    "ckczppom",
		result: 117946,
	},
}

var testPart2 = []struct {
	key    string
	result int32
}{
	{
		key:    "ckczppom",
		result: 3938038,
	},
}

func TestMine1(t *testing.T) {
	for i := range testPart1 {
		mineResult := Mine(testPart1[i].key, HashMd5, termFiveZeros)
		if testPart1[i].result != mineResult {
			t.Errorf("Number found for key %s is %d", testPart1[i].key, mineResult)
		}
	}
}

func TestMine2(t *testing.T) {
	for i := range testPart2 {
		mineResult := Mine(testPart2[i].key, HashMd5, termSixZeros)
		if testPart2[i].result != mineResult {
			t.Errorf("Number found for key %s is %d", testPart2[i].key, mineResult)
		}
	}
}
