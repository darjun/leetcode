package main

import "fmt"

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
	fmt.Println(maximumProduct([]int{-1, -2, -3}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-1, -2, -3, 4}))
	fmt.Println(maximumProductOpt1([]int{1, 2, 3}))
	fmt.Println(maximumProductOpt1([]int{-1, -2, -3}))
	fmt.Println(maximumProductOpt1([]int{1, 2, 3, 4}))
	fmt.Println(maximumProductOpt1([]int{-1, -2, -3, 4}))
}
