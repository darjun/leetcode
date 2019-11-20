package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1,1,1}, 2))
	fmt.Println(subarraySumOpt1([]int{1,1,1}, 2))
	fmt.Println(subarraySumOpt2([]int{1,1,1}, 2))
}