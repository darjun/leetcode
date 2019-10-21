package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	lo := 0
	hi := m*n - 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		i := mid / n
		j := mid % n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return false
}
