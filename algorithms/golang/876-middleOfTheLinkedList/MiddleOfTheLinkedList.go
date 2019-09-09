package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	n11 := &ListNode{Val: 1}
	n12 := &ListNode{Val: 2}
	n13 := &ListNode{Val: 3}
	n14 := &ListNode{Val: 4}
	n15 := &ListNode{Val: 5}
	n16 := &ListNode{Val: 6}
	n11.Next = n12
	n12.Next = n13
	n13.Next = n14
	n14.Next = n15
	n15.Next = n16

	fmt.Println(middleNode(n11))
}
