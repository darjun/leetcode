package main

func constructArray(n int, k int) []int {
	ans := make([]int, n, n)
	var c int
	for i := 1; i < n - k; i++ {
		ans[c] = i
		c++
	}

	for i := 0; i <= k; i++ {
		if i % 2 == 0 {
			ans[c] = n - k + i / 2
		} else {
			ans[c] = n - i / 2
		}
		c++
	}

	return ans
}
