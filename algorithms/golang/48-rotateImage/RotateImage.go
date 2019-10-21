package main

import "fmt"

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	rotate(m)
	printMatrix(m)
}
