package main

import "fmt"

func main() {
	fmt.Println(canMakePaliQueries("abcda", [][]int{[]int{3, 3, 0}, []int{1, 2, 0}, []int{0, 3, 1}, []int{0, 3, 2}, []int{0, 4, 1}}))

	fmt.Println(canMakePaliQueries("hunu", [][]int{[]int{1, 1, 1}, []int{2, 3, 0}, []int{3, 3, 1}, []int{0, 3, 2}, []int{1, 3, 3}, []int{2, 3, 1}, []int{3, 3, 1}, []int{0, 3, 0}, []int{1, 1, 1}, []int{2, 3, 0}, []int{3, 3, 1}, []int{0, 3, 1}, []int{1, 1, 1}}))
}
