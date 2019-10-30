package main

func helper(mapLetterPos map[byte]map[int]struct{}, mapUsedPos map[int]struct{}, word string, board [][]byte, index, lastPos int) bool {
	if index == len(word) {
		return true
	}

	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}

	letter := word[index]
	mapOneLetter := mapLetterPos[letter]
	if mapOneLetter == nil {
		return false
	}

	directions := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

	if lastPos == -1 {
		// 第一个字节
		for pos := range mapOneLetter {
			mapUsedPos[pos] = struct{}{}
			if helper(mapLetterPos, mapUsedPos, word, board, index+1, pos) {
				return true
			}
			delete(mapUsedPos, pos)
		}
	} else {
		for _, direction := range directions {
			lasti, lastj := lastPos>>16, lastPos&0xFFFF
			newi, newj := lasti+direction[0], lastj+direction[1]
			if newi >= 0 && newi < m && newj >= 0 && newj < n {
				if board[newi][newj] != letter {
					continue
				}
				newPos := newi<<16 + newj
				if _, isUsed := mapUsedPos[newPos]; isUsed {
					continue
				}
				mapUsedPos[newPos] = struct{}{}
				if helper(mapLetterPos, mapUsedPos, word, board, index+1, newPos) {
					return true
				}
				delete(mapUsedPos, newPos)
			}
		}
	}

	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}

	mapLetterPos := make(map[byte]map[int]struct{})
	for i, line := range board {
		for j, letter := range line {
			mapOneLetter := mapLetterPos[letter]
			if mapOneLetter == nil {
				mapOneLetter = make(map[int]struct{})
				mapLetterPos[letter] = mapOneLetter
			}

			mapOneLetter[i<<16+j] = struct{}{}
		}
	}

	mapUsedPos := make(map[int]struct{}, len(word))
	return helper(mapLetterPos, mapUsedPos, word, board, 0, -1)
}
