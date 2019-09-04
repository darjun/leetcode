package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

func main() {
	{
		n1 := &ListNode{Val: 8}
		n2 := &ListNode{Val: 4}
		n3 := &ListNode{Val: 5}
		n1.Next = n2
		n2.Next = n3

		n11 := &ListNode{Val: 4}
		n12 := &ListNode{Val: 1}
		n11.Next = n12
		n12.Next = n1

		n21 := &ListNode{Val: 5}
		n22 := &ListNode{Val: 0}
		n23 := &ListNode{Val: 1}

		n21.Next = n22
		n22.Next = n23
		n23.Next = n1
		fmt.Println(getIntersectionNodeOpt1(n11, n21).Val)
	}

	{
		n1 := &ListNode{Val: 2}
		n2 := &ListNode{Val: 4}
		n1.Next = n2

		n11 := &ListNode{Val: 0}
		n12 := &ListNode{Val: 9}
		n13 := &ListNode{Val: 1}
		n11.Next = n12
		n12.Next = n13
		n13.Next = n1

		n21 := &ListNode{Val: 3}

		n21.Next = n1
		fmt.Println(getIntersectionNodeOpt1(n11, n21).Val)
	}

	{
		n11 := &ListNode{Val: 2}
		n12 := &ListNode{Val: 4}
		n13 := &ListNode{Val: 6}
		n11.Next = n12
		n12.Next = n13

		n21 := &ListNode{Val: 1}
		n22 := &ListNode{Val: 5}
		n21.Next = n22

		fmt.Println(getIntersectionNodeOpt1(n11, n21))
	}
}
