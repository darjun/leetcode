package main

import "sort"

func helper(candidates []int, index, target int, sequence []int) [][]int {
	var ret [][]int
	for i, num := range candidates[index:] {
		if target < num {
			break
		}

		if i > 0 && candidates[index+i] == candidates[index+i-1] {
			continue
		}

		sequence = append(sequence, num)
		if target == num {
			s := make([]int, len(sequence), len(sequence))
			copy(s, sequence)
			ret = append(ret, s)
		} else {
			ret = append(ret, helper(candidates, index+i+1, target-num, sequence)...)
		}
		sequence = sequence[:len(sequence)-1]
	}

	return ret
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	sequence := make([]int, 0, 4)
	return helper(candidates, 0, target, sequence)
}
