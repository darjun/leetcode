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
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		printList(oddEvenList(n1))
	}

	{
		n1 := &ListNode{Val: 2}
		n2 := &ListNode{Val: 1}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 5}
		n5 := &ListNode{Val: 6}
		n6 := &ListNode{Val: 4}
		n7 := &ListNode{Val: 7}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		n5.Next = n6
		n6.Next = n7
		printList(oddEvenList(n1))
	}
}
