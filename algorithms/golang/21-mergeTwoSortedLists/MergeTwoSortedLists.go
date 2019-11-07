package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next

		if head != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func main() {
	{
		n11 := &ListNode{Val: 1}
		n12 := &ListNode{Val: 2}
		n13 := &ListNode{Val: 3}
		n11.Next = n12
		n12.Next = n13

		n21 := &ListNode{Val: 1}
		n22 := &ListNode{Val: 3}
		n23 := &ListNode{Val: 4}
		n21.Next = n22
		n22.Next = n23

		printList(mergeTwoLists(n11, n21))
	}
}
