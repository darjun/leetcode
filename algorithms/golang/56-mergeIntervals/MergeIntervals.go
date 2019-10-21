package main

import "fmt"

func main() {
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))

	fmt.Println(merge([][]int{[]int{1, 4}, []int{4, 5}}))
}
