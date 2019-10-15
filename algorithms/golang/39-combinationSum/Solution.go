package main

import "sort"

func helper(candidates []int, index, target int, sequence []int) [][]int {
	var ret [][]int
	for i, num := range candidates[index:] {
		if target < num {
			break
		}

		sequence = append(sequence, num)

		if target == num {
			s := make([]int, len(sequence), len(sequence))
			copy(s, sequence)
			ret = append(ret, s)
		} else {
			ret = append(ret, helper(candidates, index+i, target-num, sequence)...)
		}
		sequence = sequence[:len(sequence)-1]
	}

	return ret
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return helper(candidates, 0, target, nil)
}
