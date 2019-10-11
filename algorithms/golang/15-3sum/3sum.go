package main

import "fmt"

func main() {
	{
		fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	}

	{
		fmt.Println(threeSumOpt1([]int{-1, 0, 1, 2, -1, -4}))
	}
}
