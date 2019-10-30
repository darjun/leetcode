package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveOpt1([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveOpt2([]int{100, 4, 200, 1, 3, 2}))
}
