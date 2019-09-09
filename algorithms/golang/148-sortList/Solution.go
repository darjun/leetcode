package main

func insert(dummy, node *ListNode) {
	prev := dummy
	cur := dummy.Next
	for cur != nil && cur.Val <= node.Val {
		prev, cur = cur, cur.Next
	}

	prev.Next = node
	node.Next = cur
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	for prev, cur := head, head.Next; cur != nil; {
		if cur.Val >= prev.Val {
			prev.Next = cur
			prev, cur = cur, cur.Next
		} else {
			prev.Next = nil
			tmp := cur.Next
			insert(dummy, cur)
			cur = tmp
		}
	}

	return dummy.Next
}
