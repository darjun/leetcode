package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicateOpt1([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicateOpt1([]int{3, 1, 3, 4, 2}))
}
