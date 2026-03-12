package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		description string
		input       int
		expect      string
	}{
		{input: 1, expect: "I"},
		{input: 9, expect: "IX"},
		{input: 31, expect: "XXXI"},
		{input: 58, expect: "LVIII"},
		{input: 101, expect: "CI"},
		{input: 3749, expect: "MMMDCCXLIX"},
		{input: 1994, expect: "MCMXCIV"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.expect, func(t *testing.T) {
			if output := intToRoman(testcase.input); output != testcase.expect {
				t.Errorf("Failed to convert %d to Roman numeral %s, received: %s",
					testcase.input, testcase.expect, output)
			}
		})
	}
}
