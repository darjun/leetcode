package main

import (
	"sort"
)

func frequency(s string) int {
	var f [26]int
	for _, c := range s {
		f[c-'a']++
	}

	for _, v := range f {
		if v != 0 {
			return v
		}
	}

	return 0
}

func binarySearch(data []int, target int) int {
	i, j := 0, len(data)-1
	for i <= j {
		mid := (i + j) >> 1
		if data[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return j + 1
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordFrequency := make([]int, len(words), len(words))
	for i, word := range words {
		wordFrequency[i] = frequency(word)
	}

	sort.Ints(wordFrequency)

	numSmaller := make([]int, len(queries), len(queries))
	for i, q := range queries {
		firstLarger := binarySearch(wordFrequency, frequency(q))
		numSmaller[i] = len(wordFrequency) - firstLarger
	}

	return numSmaller
}
