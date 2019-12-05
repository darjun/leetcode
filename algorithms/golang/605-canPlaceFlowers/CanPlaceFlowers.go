package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0}, 2))
}
