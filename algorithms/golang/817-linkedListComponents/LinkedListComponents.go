package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListFromSlice(data ...int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	var head, tail *ListNode

	for _, val := range data {
		n := &ListNode{Val: val}
		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}

	return head
}

func main() {
	fmt.Println(numComponents(BuildListFromSlice(0, 1, 2, 3), []int{0, 1, 2}))

	fmt.Println(numComponents(BuildListFromSlice(0, 1, 2, 3), []int{0, 1, 3}))

	fmt.Println(numComponents(BuildListFromSlice(0, 1, 2, 3, 4), []int{0, 3, 1, 4}))
}
