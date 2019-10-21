package main

import "fmt"

func main() {
	{
		matrix := [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}
		setZeroes(matrix)
		fmt.Println(matrix)
	}

	{
		matrix := [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}
		setZeroesOpt1(matrix)
		fmt.Println(matrix)
	}
}
