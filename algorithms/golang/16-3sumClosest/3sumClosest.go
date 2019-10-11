package main

import "fmt"

func main() {
	{
		fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	}

	{
		// fmt.Println(threeSumClosestOpt1([]int{-1, 2, 1, -4}, 1))
		fmt.Println(threeSumClosestOpt1([]int{0, 2, 1, -3}, 1))
	}
}
