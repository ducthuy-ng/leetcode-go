package main

import (
	"strconv"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	testcases := []struct {
		input  [][]byte
		expect int
	}{
		{input: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, expect: 6},
		{input: [][]byte{{'0'}}, expect: 0},
		{input: [][]byte{{'1'}}, expect: 1},
	}

	for i, testcase := range testcases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if output := maximalRectangle(testcase.input); output != testcase.expect {
				t.Errorf("Failed to get maximal rectangle %d, received: %d",
					testcase.expect, output)
			}
		})
	}
}
