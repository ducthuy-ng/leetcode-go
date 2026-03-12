package main

func divide(dividend int, divisor int) int {
	isQuotientNegative := false

	if dividend < 0 && divisor > 0 {
		isQuotientNegative = true
		dividend = -dividend
	} else if dividend > 0 && divisor < 0 {
		isQuotientNegative = true
		divisor = -divisor
	} else if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	quotient := 0
	dividendCopy := dividend
	for i := 31; i >= 0; i--{
		if dividendCopy >= (divisor << i) {
			dividendCopy -= (divisor << i)
			quotient += 1 << i
		}
	}

	if isQuotientNegative {
		quotient = -quotient
	}
	if quotient < -2<<30 {
		quotient = -(2 << 30)
	} else if quotient > (2<<30)-1 {
		quotient = (2 << 30) - 1
	}
	return quotient
}
