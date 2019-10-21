package main

import "fmt"

func main() {
	{
		nums := []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)
		fmt.Println(nums)
	}

	{
		nums := []int{0, 0}
		sortColors(nums)
		fmt.Println(nums)
	}

	{
		nums := []int{1, 2, 0}
		sortColors(nums)
		fmt.Println(nums)
	}
}
