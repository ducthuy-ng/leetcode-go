package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		description string
		input       [][]byte
		expect      int
	}{
		{description: "hard", input: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, expect: 4},
		{description: "special", input: [][]byte{{'0', '1'}, {'1', '0'}}, expect: 1},
		{description: "special2", input: [][]byte{{'1'}}, expect: 1},
		{description: "minor", input: [][]byte{{'0'}}, expect: 0},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			if output := maximalSquare(testcase.input); output != testcase.expect {
				t.Errorf("Failed to calculate maximal square. Expected: %d, Received: %d", testcase.expect, output)
			}
		})
	}
}
