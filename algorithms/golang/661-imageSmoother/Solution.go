package main

func imageSmoother(M [][]int) [][]int {
	m := len(M)
	if m == 0 {
		return nil
	}

	n := len(M[0])
	if n == 0 {
		return nil
	}

	N := make([][]int, m, m)
	for i := range N {
		N[i] = make([]int, n, n)
	}

	directions := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, 1},
		[]int{1, 1},
		[]int{1, 0},
		[]int{1, -1},
		[]int{0, -1},
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := M[i][j]
			count := 1
			for _, d := range directions {
				ni := i + d[0]
				nj := j + d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					sum += M[ni][nj]
					count++
				}
			}

			N[i][j] = sum / count
		}
	}

	return N
}
