package main

func leastIntervalOpt1(tasks []byte, n int) int {
	m := make([]int, 26)
	var max int
	var maxCount int
	for _, b := range tasks {
		m[b-'A']++
		if max == m[b-'A'] {
			maxCount++
		} else if max < m[b-'A'] {
			max = m[b-'A']
			maxCount = 1
		}
	}

	partCount := max - 1
	partLength := n - (maxCount - 1)
	emptySlots := partCount * partLength
	availableTasks := len(tasks) - max*maxCount
	idles := emptySlots - availableTasks
	if idles < 0 {
		idles = 0
	}

	return len(tasks) + idles
}
