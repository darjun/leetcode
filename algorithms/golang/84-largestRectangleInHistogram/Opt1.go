package main

func getMinIndex(heights []int, l, r int) int {
	minIndex := -1

	for i := l; i <= r; i++ {
		if minIndex == -1 || heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}

func helper(heights []int, l, r int) int {
	if l == r {
		return heights[l]
	}

	if l > r {
		return 0
	}

	minIndex := getMinIndex(heights, l, r)
	leftMax := helper(heights, l, minIndex-1)
	rightMax := helper(heights, minIndex+1, r)

	maxArea := heights[minIndex] * (r - l + 1)
	if maxArea < leftMax {
		maxArea = leftMax
	}

	if maxArea < rightMax {
		maxArea = rightMax
	}

	return maxArea
}

func largestRectangleAreaOpt1(heights []int) int {
	return helper(heights, 0, len(heights)-1)
}
