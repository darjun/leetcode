package main

func insert(intervals [][]int, newInterval []int) [][]int {
	lo := 0
	hi := len(intervals) - 1
	foundIndex := -1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if intervals[mid][0] == newInterval[0] {
			foundIndex = mid
			break
		} else if intervals[mid][0] < newInterval[0] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	var ans [][]int
	var start int
	if foundIndex == -1 {
		start = hi
		if hi < 0 {
			start = 0
		}
	} else {
		start = foundIndex
	}

	if start > 0 {
		ans = intervals[0 : start-1]
	}

	i := start
	for ; i < len(intervals); i++ {
		if intervals[i][0] <= newInterval[0] && intervals[i][1] >= newInterval[0] {
			newInterval[0] = intervals[i][0]

			if intervals[i][1] > newInterval[1] {
				newInterval[1] = intervals[i][1]
			}
		}
	}

	ans = append(ans, newInterval)
	ans = append(ans, intervals[i:]...)

	return ans
}
