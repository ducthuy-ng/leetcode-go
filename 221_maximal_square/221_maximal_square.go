package main

func maximalSquare(matrix [][]byte) int {
	colCount := len(matrix[0])

	prevMaxSquareLen := make([]int, colCount)
	maxArea := 0

	for _, row := range matrix {
		maxSquareLenToCurrent := make([]int, colCount)
		if row[0] == '1' {
			maxSquareLenToCurrent[0] = 1
			if maxArea < 1 {
				maxArea = 1
			}
		}

		for j := 1; j < colCount; j++ {
			if row[j] == '0' {
				continue
			}

			prevMinLength := 1000
			prevMinLength = min(prevMinLength, prevMaxSquareLen[j], prevMaxSquareLen[j-1], maxSquareLenToCurrent[j-1])

			maxSquareLenToCurrent[j] = prevMinLength + 1
			currentArea := (prevMinLength + 1) * (prevMinLength + 1)
			if currentArea > maxArea {
				maxArea = currentArea
			}
		}
		prevMaxSquareLen = maxSquareLenToCurrent
	}
	return maxArea
}
