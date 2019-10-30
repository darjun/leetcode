package main

func largestRectangleArea(heights []int) int {
	var maxArea int
	l := len(heights)
	for i, height := range heights {
		left := i
		right := i
		for left > 0 && heights[left-1] >= height {
			left--
		}

		for right < l-1 && heights[right+1] >= height {
			right++
		}

		area := height * (right - left + 1)
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
