package main

func maxArea(height []int) int {
	max := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			var area int
			if height[i] > height[j] {
				area = height[j] * (j - i)
			} else {
				area = height[i] * (j - i)
			}

			if max < area {
				max = area
			}
		}
	}

	return max
}
