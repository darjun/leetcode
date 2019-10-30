package main

func cloneSlice(nums []int) []int {
	r := make([]int, len(nums), len(nums))
	for i, num := range nums {
		r[i] = num
	}

	return r
}

func helper(start, k, n int, path []int, ans *[][]int) {
	if k == 0 {
		if n == 0 {
			*ans = append(*ans, cloneSlice(path))
		}
		return
	}

	for (start+start+k-1)*k/2 <= n && start <= 9 {
		path = append(path, start)
		helper(start+1, k-1, n-start, path, ans)
		start++
		path = path[:len(path)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	path := make([]int, 0, 1)
	helper(1, k, n, path, &ans)
	return ans
}
