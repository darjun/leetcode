package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-1, -2, -9, -6}))
	fmt.Println(maxProduct([]int{-2}))

	fmt.Println(maxProductOpt1([]int{2, 3, -2, 4}))
	fmt.Println(maxProductOpt1([]int{-2, 0, -1}))
	fmt.Println(maxProductOpt1([]int{-1, -2, -9, -6}))
	fmt.Println(maxProductOpt1([]int{-2}))
}
