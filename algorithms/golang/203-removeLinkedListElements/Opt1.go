package main

func removeElementsOpt1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return head
	}

	l := head
	for l.Next != nil {
		if l.Next.Val == val {
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}

	return head
}
