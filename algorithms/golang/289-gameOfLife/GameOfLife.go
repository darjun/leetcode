package main

import "fmt"

func main() {
	{
		board := [][]int{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
			[]int{0, 0, 0},
		}
		gameOfLife(board)
		fmt.Println(board)
	}
}
