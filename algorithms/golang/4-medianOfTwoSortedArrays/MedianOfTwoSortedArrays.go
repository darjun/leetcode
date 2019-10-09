package main

import "fmt"

func main() {
	{
		fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	}

	{
		fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	}

	{
		fmt.Println(findMedianSortedArraysOpt1([]int{1, 3}, []int{2}))
	}

	{
		fmt.Println(findMedianSortedArraysOpt1([]int{1, 3}, []int{2, 4}))
	}
}
