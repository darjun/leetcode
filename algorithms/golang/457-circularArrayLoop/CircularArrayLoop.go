package main

import "fmt"

func main() {
	fmt.Println(circularArrayLoop([]int{2,-1,1,2,2}))
	fmt.Println(circularArrayLoop([]int{-1,2}))
	fmt.Println(circularArrayLoop([]int{-2,1,-1,-2,-2}))
}