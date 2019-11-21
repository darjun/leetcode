package main

import "fmt"

func main() {
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 2, 4))
}