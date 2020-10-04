//+build integration

package main_test

import (
	"testing"
)

func TestSomeIntegrationTestCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 4},
		{input: -1, expected: 1},
		{input: 0, expected: 2}, //
		{input: -5, expected: -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("Test Failed: input=[%v], expected=[%v], output=[%v]", test.input, test.expected, output)
		}
	}
}
