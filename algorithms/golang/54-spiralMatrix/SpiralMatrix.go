package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))

	fmt.Println(spiralOrderOpt1([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
}
