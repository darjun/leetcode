package main

import "fmt"

func main() {
	{
		fmt.Println(findPairs([]int{3,1,4,1,5}, 2))
		fmt.Println(findPairs([]int{1,2,3,4,5}, 1))
		fmt.Println(findPairs([]int{1,3,1,5,4}, 0))
	}

	{
		fmt.Println(findPairsOpt1([]int{3,1,4,1,5}, 2))
		fmt.Println(findPairsOpt1([]int{1,2,3,4,5}, 1))
		fmt.Println(findPairsOpt1([]int{1,3,1,5,4}, 0))
	}
}