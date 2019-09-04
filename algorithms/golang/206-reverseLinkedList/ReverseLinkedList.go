package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *ListNode = nil, head, head.Next
	for cur != nil {
		cur.Next = prev
		prev, cur = cur, next
		if next != nil {
			next = next.Next
		}
	}

	return prev
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
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		printList(reverse(n1))
	}

	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		printList(reverse2(n1))
	}
}
