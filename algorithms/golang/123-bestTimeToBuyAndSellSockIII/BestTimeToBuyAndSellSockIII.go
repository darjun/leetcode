package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 5, 3, 1}))

	fmt.Println(maxProfitOpt1([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfitOpt1([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfitOpt1([]int{7, 6, 5, 3, 1}))
}
