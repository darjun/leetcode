package main

func subsets(nums []int) [][]int {
	n := len(nums)
	p := 1 << uint(n)
	ans := make([][]int, 0, p)
	ans = append(ans, []int{})
	for i := 1; i < p; i++ {
		sub := make([]int, 0, 1)
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) != 0 {
				sub = append(sub, nums[j])
			}
		}

		ans = append(ans, sub)
	}
	return ans
}
