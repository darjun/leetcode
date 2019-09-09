package main

// merge sort
func sortListOpt1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var preSlow *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil
	return merge(sortListOpt1(head), sortListOpt1(slow))
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, tail *ListNode
	p1 := l1
	p2 := l2
	if p1.Val > p2.Val {
		head = p2
		tail = p2
		p2 = p2.Next
	} else {
		head = p1
		tail = p1
		p1 = p1.Next
	}

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			tail.Next = p2
			tail = p2
			p2 = p2.Next
		} else {
			tail.Next = p1
			tail = p1
			p1 = p1.Next
		}
	}

	for p1 != nil {
		tail.Next = p1
		tail = p1
		p1 = p1.Next
	}

	for p2 != nil {
		tail.Next = p2
		tail = p2
		p2 = p2.Next
	}

	tail.Next = nil
	return head
}
