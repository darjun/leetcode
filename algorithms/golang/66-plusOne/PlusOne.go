package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))

	fmt.Println(plusOne([]int{4, 3, 2, 1}))

	fmt.Println(plusOne([]int{9, 9, 9, 9}))

	fmt.Println(plusOne([]int{}))
}
