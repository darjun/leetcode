package main

import "fmt"

func main() {
	fmt.Println(imageSmoother([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}))
	fmt.Println(imageSmoother([][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}))

	fmt.Println(imageSmootherOpt1([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}))
	fmt.Println(imageSmootherOpt1([][]int{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}))
}
