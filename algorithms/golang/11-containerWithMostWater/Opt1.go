package main

func maxAreaOpt1(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	i := 0
	j := len(height) - 1

	var max int
	for i < j {
		var area int
		if height[j] > height[i] {
			area = height[i] * (j - i)
		} else {
			area = height[j] * (j - i)
		}

		if max < area {
			max = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
