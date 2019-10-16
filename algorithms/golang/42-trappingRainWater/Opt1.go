package main

func trapOpt1(height []int) int {
	l := len(height)
	leftMax := make([]int, 0, len(height))
	rightMax := make([]int, 0, len(height))

	max := 0
	for _, h := range height {
		if h > max {
			max = h
		}

		leftMax = append(leftMax, max)
	}

	max = 0
	for i := l - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}

		rightMax = append(rightMax, max)
	}

	var water int
	for i := 1; i < l-1; i++ {
		if leftMax[i] < rightMax[l-i-1] {
			water += leftMax[i] - height[i]
		} else {
			water += rightMax[l-i-1] - height[i]
		}
	}

	return water
}
