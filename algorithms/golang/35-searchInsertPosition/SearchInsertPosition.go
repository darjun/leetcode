package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
