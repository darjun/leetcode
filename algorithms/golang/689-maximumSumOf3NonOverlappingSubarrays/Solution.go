package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	w := make([]int, len(nums)-k+1)
	sum := 0
	for i, num := range nums {
		sum += num
		if i >= k {
			sum -= nums[i-k]
		}

		if i >= k-1 {
			w[i-k+1] = sum
		}
	}

	left := make([]int, len(w))
	var best int
	for i := range w {
		if w[i] > w[best] {
			best = i
		}

		left[i] = best
	}

	right := make([]int, len(w))
	best = len(w) - 1
	for i := len(w) - 1; i >= 0; i-- {
		if w[i] >= w[best] {
			best = i
		}
		right[i] = best
	}

	ans := []int{-1, -1, -1}
	for j := k; j < len(w)-k; j++ {
		i := left[j-k]
		k := right[j+k]

		if ans[0] == -1 || w[i]+w[j]+w[k] > w[ans[0]]+w[ans[1]]+w[ans[2]] {
			ans[0] = i
			ans[1] = j
			ans[2] = k
		}
	}

	return ans
}
