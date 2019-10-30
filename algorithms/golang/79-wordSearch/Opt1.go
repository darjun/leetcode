package main

func hasPath(board [][]byte, word string, m, n, index, lasti, lastj int) bool {
	if index == len(word) {
		return true
	}

	directions := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	for _, direction := range directions {
		newi, newj := lasti+direction[0], lastj+direction[1]
		if newi >= 0 && newi < m && newj >= 0 && newj < n {
			if board[newi][newj] != word[index] {
				continue
			}
			board[newi][newj] = 0
			if hasPath(board, word, m, n, index+1, newi, newj) {
				return true
			}
			board[newi][newj] = word[index]
		}
	}

	return false
}

func existOpt1(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	for i, line := range board {
		for j, letter := range line {
			if letter == word[0] {
				board[i][j] = 0
				if hasPath(board, word, m, n, 1, i, j) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}

	return false
}
