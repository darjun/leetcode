package main

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1, rowIndex+1)

	// C(n,m+1) = C(n,m) * (n-m)/(m+1)
	ans[0] = 1
	last := 1
	for i := 1; i <= rowIndex; i++ {
		last = last * (rowIndex - i + 1) / i
		ans[i] = last
	}

	return ans
}
