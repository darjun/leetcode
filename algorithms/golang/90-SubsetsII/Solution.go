package main

func cloneSlice(nums []int) []int {
	ret := make([]int, len(nums))
	copy(ret, nums)
	return ret
}

func helper(counts [][2]int, index int, subset []int, ans *[][]int) {
	if index == len(counts) {
		*ans = append(*ans, cloneSlice(subset))
		return
	}

	for c := 0; c <= counts[index][1]; c++ {
		oldLen := len(subset)
		for j := 0; j < c; j++ {
			subset = append(subset, counts[index][0])
		}
		helper(counts, index+1, subset, ans)
		subset = subset[:oldLen]
	}
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0, len(nums))

	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	counts := make([][2]int, 0, len(countMap))
	for num, count := range countMap {
		counts = append(counts, [2]int{num, count})
	}

	helper(counts, 0, make([]int, 0, 1), &ans)
	return ans
}
