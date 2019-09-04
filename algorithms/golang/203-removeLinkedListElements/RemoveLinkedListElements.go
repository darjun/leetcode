package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	prev, cur := head, head.Next
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev, cur = cur, cur.Next
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
	{
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 6}
		n4 := &ListNode{Val: 3}
		n5 := &ListNode{Val: 4}
		n6 := &ListNode{Val: 5}
		n7 := &ListNode{Val: 6}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		n6.Next = n7
		printList(removeElementsOpt1(n1, 6))
	}
}
