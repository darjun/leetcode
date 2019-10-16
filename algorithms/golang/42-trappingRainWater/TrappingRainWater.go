package main

import "fmt"

func main() {
	{
		fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

		fmt.Println(trap([]int{2, 0, 2}))
	}

	{
		fmt.Println(trapOpt1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

		fmt.Println(trapOpt1([]int{2, 0, 2}))
	}
}
