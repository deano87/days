package main

import "testing"

var deliveryList = []struct {
	Input          string
	ExpectedResult int
}{
	{
		Input:          "^v",
		ExpectedResult: 3,
	},
	{
		Input:          "^v^v^v^v^v",
		ExpectedResult: 11,
	},
}

// TestDelivery test that the number of houses that got presents is corrent
func TestDelivery(t *testing.T) {
	for i := range deliveryList {
		parseResult := ParseElfInstructions(deliveryList[i].Input)
		if parseResult != deliveryList[i].ExpectedResult {
			t.Errorf("Expected result (%d) doesn't match the number of houses returned (%d)", deliveryList[i].ExpectedResult, parseResult)
		}
	}
}
