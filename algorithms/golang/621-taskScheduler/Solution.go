package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	for _, b := range tasks {
		m[b-'A']++
	}

	sort.Ints(m)
	var count int
	for m[25] > 0 {
		var i int
		for i <= n {
			if m[25] == 0 {
				break
			}
			if i < 26 && m[25-i] > 0 {
				m[25-i]--
			}
			i++
			count++
		}

		sort.Ints(m)
	}

	return count
}
