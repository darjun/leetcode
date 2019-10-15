package main

import "sort"

type Item struct {
	index    int
	target   int
	sequence []int
}

func combinationSumOpt1(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int
	stack := make([]Item, 0, 1)
	stack = append(stack, Item{0, target, nil})

	l := len(stack)
	for l > 0 {
		item := stack[l-1]
		stack = stack[:l-1]
		index, t, sequnce := item.index, item.target, item.sequence

		for i, num := range candidates[index:] {
			if t < num {
				break
			}

			s := make([]int, len(sequnce), len(sequnce))
			copy(s, sequnce)
			s = append(s, num)
			if num < t {
				stack = append(stack, Item{index + i, t - num, s})
			} else {
				ans = append(ans, s)
			}
		}

		l = len(stack)
	}
	return ans
}
