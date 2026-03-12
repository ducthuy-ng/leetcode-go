package main

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	testcases := []struct {
		input  []int
		expect int
	}{
		{input: []int{2, 1, 5, 6, 2, 3}, expect: 10},
		{input: []int{2, 4}, expect: 4},
		{input: []int{1, 2, 3, 4, 5}, expect: 9},
		{input: []int{1}, expect: 1},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("%v", testcase.input), func(t *testing.T) {
			if output := largestRectangleArea(testcase.input); output != testcase.expect {
				t.Errorf("Failed to get largest rectangle area. Testcase: %v, expect: %v, received: %d",
					testcase.input, testcase.expect, output)
			}
		})
	}
}
