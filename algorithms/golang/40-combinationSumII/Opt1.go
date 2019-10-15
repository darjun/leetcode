package main

import "sort"

type Item struct {
	index    int
	target   int
	sequence []int
}

func combinationSum2Opt1(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int

	stack := make([]Item, 0, 1)
	stack = append(stack, Item{0, target, nil})

	l := len(stack)
	for l > 0 {
		item := stack[l-1]
		stack = stack[:l-1]

		index, t, sequence := item.index, item.target, item.sequence

		for i, num := range candidates[index:] {
			if t < num {
				break
			}

			if i > 0 && candidates[index+i] == candidates[index+i-1] {
				continue
			}

			s := make([]int, len(sequence), len(sequence))
			copy(s, sequence)
			s = append(s, num)
			if t == num {
				ans = append(ans, s)
				break
			} else {
				stack = append(stack, Item{index + i + 1, t - num, s})
			}
		}

		l = len(stack)
	}

	return ans
}
