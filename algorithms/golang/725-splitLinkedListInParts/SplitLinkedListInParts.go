package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListFromSlice(data []int) *ListNode {
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

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		fmt.Print(" -> ")
		head = head.Next
	}

	fmt.Println("NULL")
}

func main() {
	for _, l := range splitListToParts(BuildListFromSlice([]int{1, 2, 3}), 5) {
		printList(l)
	}

	fmt.Println("========================")

	for _, l := range splitListToParts(BuildListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3) {
		printList(l)
	}
}
