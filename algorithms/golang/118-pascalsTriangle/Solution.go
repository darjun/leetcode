package main

func cloneSlice(nums []int) []int {
	ret := make([]int, len(nums), len(nums))
	copy(ret, nums)
	return ret
}

func generate(numRows int) [][]int {
	ans := make([][]int, 0, 1)
	if numRows >= 1 {
		ans = append(ans, []int{1})
	}

	if numRows >= 2 {
		ans = append(ans, []int{1, 1})
	}

	if numRows <= 2 {
		return ans
	}

	for i := 3; i <= numRows; i++ {
		lastRow := ans[len(ans)-1]
		nowRow := make([]int, len(lastRow), len(lastRow))
		nowRow[0] = 1
		nowRow = append(nowRow, 1)
		for i := 1; i < len(nowRow)-1; i++ {
			nowRow[i] = lastRow[i-1] + lastRow[i]
		}
		ans = append(ans, nowRow)
	}

	return ans
}
