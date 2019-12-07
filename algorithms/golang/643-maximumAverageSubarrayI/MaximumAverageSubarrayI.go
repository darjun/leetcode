package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))

	fmt.Println(findMaxAverageOpt1([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverageOpt1([]int{5}, 1))
}
