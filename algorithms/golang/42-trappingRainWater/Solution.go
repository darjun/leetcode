package main

func trap(height []int) int {
	l := len(height)
	var water int
	for i := 1; i < l-1; i++ {
		var maxLeft, maxRight int
		for j := i; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}

		for j := i; j < l; j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}

		if maxLeft < maxRight {
			water += maxLeft - height[i]
		} else {
			water += maxRight - height[i]
		}
	}

	return water
}
