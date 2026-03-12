package main

type LargestAreaDetails struct {
	Width  int
	Height int
}

func largestRectangleArea(heights []int) int {
	type stackItem struct {
		Index  int
		Height int
	}
	stack := make([]stackItem, 0, len(heights))
	maxArea := 0

	for i, height := range heights {
		if len(stack) == 0 || stack[len(stack)-1].Height <= height {
			stack = append(stack, stackItem{Index: i, Height: height})
			continue
		}

		for {
			if len(stack) == 0 {
				stack = append(stack, stackItem{Index: 0, Height: height})
				break
			}

			stackTop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if stackTop.Height > height {
				stackTopArea := (i - stackTop.Index) * stackTop.Height
				if stackTopArea > maxArea {
					maxArea = stackTopArea
				}
			}

			if len(stack) > 0 && stack[len(stack)-1].Height <= height {
				stack = append(stack, stackItem{Index: stackTop.Index, Height: height})
				break
			}

		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if maxArea < stack[i].Height*(len(heights)-stack[i].Index) {
			maxArea = stack[i].Height * (len(heights) - stack[i].Index)
		}
	}

	return maxArea
}

func largestRectangleAreaAttemp1(heights []int) int {
	maxRectArea := make([]LargestAreaDetails, len(heights))
	maxRectArea[0].Width = 1
	maxRectArea[0].Height = heights[0]

	for i := 1; i < len(heights); i++ {
		currentMaxArea := LargestAreaDetails{Width: 1, Height: heights[i]}

		incrementHeight := min(maxRectArea[i-1].Height, heights[i])

		if incrementHeight*(maxRectArea[i-1].Width+1) > currentMaxArea.Width*currentMaxArea.Height {
			currentMaxArea.Height = incrementHeight
			currentMaxArea.Width = maxRectArea[i-1].Width + 1
		}

		maxRectArea[i] = currentMaxArea
	}

	maxArea := 0
	for _, rect := range maxRectArea {
		if maxArea < rect.Width*rect.Height {
			maxArea = rect.Width * rect.Height
		}
	}
	return maxArea
}
