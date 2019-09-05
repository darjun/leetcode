package main

// 反转列表
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

// 先将列表拆成前后两个部分，然后反转第二个部分，再将第二个列表插入第一个之中。
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prevSlow := (*ListNode)(nil)
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		// 偶数个节点
	} else {
		// 奇数个节点
		prevSlow = slow
		slow = slow.Next
	}
	prevSlow.Next = nil

	// 反转第二部分
	second := reverse(slow)

	// 插入第一部分
	cur := head
	next := head.Next
	for cur != nil && second != nil {
		secondNext := second.Next
		cur.Next = second
		second.Next = next

		cur = next
		if next != nil {
			next = next.Next
		}

		second = secondNext
	}

	return head
}
