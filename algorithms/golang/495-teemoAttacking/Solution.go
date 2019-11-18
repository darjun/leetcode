package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}

	var total int
	for i:=0;i<n-1;i++ {
		diff := timeSeries[i+1] - timeSeries[i]
		if diff > duration {
			total += duration
		} else {
			total += diff
		}
	}

	return total + duration
}