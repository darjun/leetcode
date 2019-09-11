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
	printList(removeZeroSumSublistsOpt1(BuildListFromSlice([]int{1, 2, -3, 3, 1})))

	printList(removeZeroSumSublistsOpt1(BuildListFromSlice([]int{1, 2, 3, -3, 4})))

	printList(removeZeroSumSublistsOpt1(BuildListFromSlice([]int{1, 2, 3, -3, -2})))

	printList(removeZeroSumSublistsOpt1(BuildListFromSlice([]int{-1, 1, 0, 1})))

	printList(removeZeroSumSublistsOpt1(BuildListFromSlice([]int{2, 2, -2, 1, -1, -1})))
}
