package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// node.Next的值复制到node中，然后删掉node.Next
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
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
		n1 := &ListNode{Val: 4}
		n2 := &ListNode{Val: 5}
		n3 := &ListNode{Val: 1}
		n4 := &ListNode{Val: 9}
		n1.Next = n2
		n2.Next = n3
		n3.Next = n4

		deleteNode(n2)
		printList(n1)
	}
}
