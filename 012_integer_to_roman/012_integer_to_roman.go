package main

import (
	"math"
	"strconv"
	"strings"
)

var SYMBOLS = map[int]rune{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}

func intToRoman(num int) string {
	strNum := []rune(strconv.Itoa(num))
	decimal := int(math.Pow10(len(strNum)))

	var result strings.Builder
	for _, char := range strNum {
		decimal /= 10
		if char == '0' {
			continue
		}
		if char < '4' {
			symbol := SYMBOLS[decimal]
			for range char - '0' {
				result.WriteRune(symbol)
			}
		} else if char == '4' {
			previousSymbol := SYMBOLS[decimal]
			symbol := SYMBOLS[decimal*5]
			result.WriteRune(previousSymbol)
			result.WriteRune(symbol)
		} else if char < '9' {
			symbol := SYMBOLS[decimal*5]
			result.WriteRune(symbol)

			previousSymbol := SYMBOLS[decimal]
			for range (char - '0') - 5 {
				result.WriteRune(previousSymbol)
			}
		} else {
			previousSymbol := SYMBOLS[decimal]
			symbol := SYMBOLS[decimal*10]
			result.WriteRune(previousSymbol)
			result.WriteRune(symbol)
		}
	}
	return result.String()
}
