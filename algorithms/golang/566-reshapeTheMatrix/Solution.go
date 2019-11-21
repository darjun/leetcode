package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
    m := len(nums)
    if m == 0 {
        return nil
    }
    
	n := len(nums[0])
	if n == 0 {
		return nil
	}

	if m * n != r * c {
		return nums
	}

	if m == r && n == c {
		return nums
	}

	ans := make([][]int, r, r)
	for i := range ans {
		ans[i] = make([]int, c, c)
	}

	var newi, newj int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[newi][newj] = nums[i][j]
			newj++
			if newj == c {
				newi++
				newj = 0
			}
		}
	}

	return ans
}