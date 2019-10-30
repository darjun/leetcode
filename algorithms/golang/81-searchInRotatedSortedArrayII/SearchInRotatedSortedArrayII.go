package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{0, 0, 1, 2, 0, 0, 0}, 0))
	fmt.Println(search([]int{0, 0, 0, 0, 0, 0, 0}, 0))
}
