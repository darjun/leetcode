package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
	{
		n11 := &ListNode{Val: 7}
		n12 := &ListNode{Val: 2}
		n13 := &ListNode{Val: 4}
		n14 := &ListNode{Val: 3}
		n11.Next = n12
		n12.Next = n13
		n13.Next = n14

		n21 := &ListNode{Val: 5}
		n22 := &ListNode{Val: 6}
		n23 := &ListNode{Val: 4}
		n21.Next = n22
		n22.Next = n23
		printList(addTwoNumbers(n11, n21))
	}

	{
		n11 := &ListNode{Val: 5}

		n21 := &ListNode{Val: 5}
		printList(addTwoNumbers(n11, n21))
	}
}
