package main

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	toDie := make(map[int]struct{})
	toReproduce := make(map[int]struct{})

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

	for i, line := range board {
		for j, value := range line {
			var liveNeighbors int
			for _, direction := range directions {
				newi := i + direction[0]
				newj := j + direction[1]
				if newi >= 0 && newi < m && newj >= 0 && newj < n {
					if board[newi][newj] == 1 {
						liveNeighbors++
					}
				}
			}
			if value == 0 && liveNeighbors == 3 {
				toReproduce[i<<16+j] = struct{}{}
			} else if value == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				toDie[i<<16+j] = struct{}{}
			}
		}
	}

	for die := range toDie {
		i := die >> 16
		j := die & 0xFFFF
		board[i][j] = 0
	}

	for reproduce := range toReproduce {
		i := reproduce >> 16
		j := reproduce & 0xFFFF
		board[i][j] = 1
	}
}
