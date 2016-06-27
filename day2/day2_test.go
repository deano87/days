package main

import "testing"

var testList = struct {
	presents                  PresentList
	expectedWrapperDimensions int
	expectedRibbonLength      int
}{
	presents: PresentList{
		Present{Width: 1, Height: 1, Length: 1},
		Present{Width: 2, Height: 1, Length: 1},
	},
	expectedWrapperDimensions: 18,
	expectedRibbonLength:      11,
}

// TestWrappingDimensions tests wrapping paper calculation
func TestWrappingDimensions(t *testing.T) {
	if testList.presents.GetTotalWrapperDimensions() != testList.expectedWrapperDimensions {
		t.Error("Total wrapping dimensions don't match the expected result")
	}
}

// TestRibbonLength tests ribbon length calculation
func TestRibbonLength(t *testing.T) {
	if testList.presents.GetTotalRibbonLength() != testList.expectedRibbonLength {
		t.Error("Total ribbon length doesn't match the expected result")
	}
}
