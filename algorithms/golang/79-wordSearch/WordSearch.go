package main

import "fmt"

func main() {
	{
		board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
		fmt.Println(exist(board, "ABCCED"))
		fmt.Println(exist(board, "SEE"))
		fmt.Println(exist(board, "ABCB"))
	}

	{
		board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
		fmt.Println(existOpt1(board, "ABCCED"))
		fmt.Println(existOpt1(board, "SEE"))
		fmt.Println(existOpt1(board, "ABCB"))
	}
}
