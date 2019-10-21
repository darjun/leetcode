package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}

		return intervals[i][1] < intervals[j][1]
	})

	curInterval := intervals[0]
	var ans [][]int
	for _, interval := range intervals[1:] {
		if interval[0] >= curInterval[0] && interval[0] <= curInterval[1] {
			if curInterval[1] < interval[1] {
				curInterval[1] = interval[1]
			}
		} else {
			ans = append(ans, curInterval)
			curInterval = interval
		}
	}

	ans = append(ans, curInterval)
	return ans
}
