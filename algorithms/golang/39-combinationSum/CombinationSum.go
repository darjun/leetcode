package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))

	fmt.Println(combinationSum([]int{2, 3, 5}, 8))

	fmt.Println(combinationSumOpt1([]int{2, 3, 6, 7}, 7))

	fmt.Println(combinationSumOpt1([]int{2, 3, 5}, 8))
}
