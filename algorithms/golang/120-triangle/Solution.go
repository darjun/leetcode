package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	lastLine := triangle[len(triangle)-1]
	min := make([]int, len(lastLine), len(lastLine))
	copy(min, lastLine)

	for i := len(triangle) - 2; i >= 0; i-- {
		line := triangle[i]
		for j := 0; j < len(line); j++ {
			if min[j] > min[j+1] {
				min[j] = line[j] + min[j+1]
			} else {
				min[j] = line[j] + min[j]
			}
		}
	}

	return min[0]
}
