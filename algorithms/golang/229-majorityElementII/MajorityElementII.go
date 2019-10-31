package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))

	fmt.Println(majorityElementOpt1([]int{3, 2, 3}))
	fmt.Println(majorityElementOpt1([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
